package transport

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"test/log"
	"test/pkg/sip"
)

// Transport error
type Error interface {
	net.Error
	// Network indicates network level errors
	Network() bool
}

type RequestCallback func(sip.RequestMethod, sip.Message)

// Layer serves client and server transactions.
type TransportLayer interface {
	Cancel()
	Done() <-chan struct{}
	Messages() <-chan sip.Message
	Errors() <-chan error
	// Listen starts listening on `addr` for each registered protocol.
	Listen(network string, addr string, options ...ListenOption) error
	// Send sends message on suitable protocol.
	Send(msg sip.Message) error
	String() string
	IsReliable(network string) bool
	IsStreamed(network string) bool
}

type transportLayer struct {
	protocols   *protocolStore
	listenPorts map[string][]sip.Port
	ip          net.IP
	dnsResolver *net.Resolver
	msgMapper   sip.MessageMapper

	msgs     chan sip.Message
	errs     chan error
	pmsgs    chan sip.Message
	perrs    chan error
	canceled chan struct{}
	done     chan struct{}

	wg         sync.WaitGroup
	cancelOnce sync.Once

	log log.Logger
}

var protocolFactory ProtocolFactory = func(
	network string,
	output chan<- sip.Message,
	errs chan<- error,
	cancel <-chan struct{},
	msgMapper sip.MessageMapper,
	logger log.Logger,
) (Protocol, error) {
	switch strings.ToLower(network) {
	case "udp":
		return NewUdpProtocol(output, errs, cancel, msgMapper, logger), nil
	case "tcp":
		return NewTcpProtocol(output, errs, cancel, msgMapper, logger), nil
	case "tls":
		return nil, UnsupportedProtocolError(fmt.Sprintf("protocol %s is not supported", network))

		// return NewTlsProtocol(output, errs, cancel, msgMapper, logger), nil
	case "ws":
		return nil, UnsupportedProtocolError(fmt.Sprintf("protocol %s is not supported", network))

		// return NewWsProtocol(output, errs, cancel, msgMapper, logger), nil
	case "wss":
		return nil, UnsupportedProtocolError(fmt.Sprintf("protocol %s is not supported", network))

		// return NewWssProtocol(output, errs, cancel, msgMapper, logger), nil
	default:
		return nil, UnsupportedProtocolError(fmt.Sprintf("protocol %s is not supported", network))
	}
}

// 生成TransportLayer对象
func NewTransportLayer(
	ip net.IP,
	dnsResolver *net.Resolver,
	msgMapper sip.MessageMapper,
	logger log.Logger,
) TransportLayer {
	// TransportLayer初始化
	tpl := &transportLayer{
		protocols:   newProtocolStore(),
		listenPorts: make(map[string][]sip.Port),
		ip:          ip,
		dnsResolver: dnsResolver,
		msgMapper:   msgMapper,

		msgs:     make(chan sip.Message),
		errs:     make(chan error),
		pmsgs:    make(chan sip.Message),
		perrs:    make(chan error),
		canceled: make(chan struct{}),
		done:     make(chan struct{}),
	}
	tpl.log = logger.
		WithPrefix("transport.Layer").
		WithFields(map[string]interface{}{
			"transport_layer_ptr": fmt.Sprintf("%p", tpl),
		})

	go tpl.serveProtocols()

	return tpl
}

// 开始监听
func (tpl *transportLayer) Listen(network string, addr string, options ...ListenOption) error {
	protocol, ok := tpl.protocols.get(protocolKey(network))

	if !ok {
		var err error
		protocol, err = protocolFactory(
			network,
			tpl.pmsgs,
			tpl.perrs,
			tpl.canceled,
			tpl.msgMapper,
			tpl.Log(),
		)
		if err != nil {
			return err
		}
		// 塞入线程池
		tpl.protocols.put(protocolKey(protocol.Network()), protocol)
	}

	target, err := NewTargetFromAddr(addr)
	if err != nil {
		return err
	}
	// 启动对应协议的端口监听
	err = protocol.Listen(target, options...)

	if err == nil {
		if _, ok := tpl.listenPorts[protocol.Network()]; !ok {
			if tpl.listenPorts[protocol.Network()] == nil {
				tpl.listenPorts[protocol.Network()] = make([]sip.Port, 0)
			}
			tpl.listenPorts[protocol.Network()] = append(tpl.listenPorts[protocol.Network()], *target.Port)
		}
	}

	return nil
}

// 启动协程 监听消息进入
func (tpl *transportLayer) serveProtocols() {
	defer func() {
		tpl.dispose()
		close(tpl.done)
	}()

	tpl.Log().Debug("begin serve protocols")
	defer tpl.Log().Debug("stop serve protocols")

	for {
		select {
		case <-tpl.canceled:
			return
		case msg := <-tpl.pmsgs:
			tpl.handleMessage(msg)
		case err := <-tpl.perrs:
			tpl.handlerError(err)
		}
	}
}

// 消息处理机制
func (tpl *transportLayer) handleMessage(msg sip.Message) {
	logger := tpl.Log().WithFields(msg.Fields())

	logger.Debugf("接受SIP消息:\n%s", msg)

	// pass up message
	select {
	case <-tpl.canceled:
	case tpl.msgs <- msg:
		logger.Trace("SIP message passed up")
	}
}

// 错误机制
func (tpl *transportLayer) handlerError(err error) {
	// TODO: implement re-connection strategy for listeners
	var terr Error
	if errors.As(err, &terr) {
		// currently log
		tpl.Log().Warnf("SIP transport error: %s", err)
	}

	logger := tpl.Log().WithFields(log.Fields{
		"sip_error": err.Error(),
	})

	logger.Trace("passing up error...")

	select {
	case <-tpl.canceled:
	case tpl.errs <- err:
		logger.Trace("error passed up")
	}
}

func (tpl *transportLayer) Send(msg sip.Message) error {

	return nil
}

// 销毁
func (tpl *transportLayer) dispose() {
	tpl.Log().Debug("disposing...")
	// wait for protocols
	for _, protocol := range tpl.protocols.all() {
		tpl.protocols.drop(protocolKey(protocol.Network()))
		<-protocol.Done()
	}

	tpl.listenPorts = make(map[string][]sip.Port)

	close(tpl.pmsgs)
	close(tpl.perrs)
	close(tpl.msgs)
	close(tpl.errs)
}

// 取消
func (tpl *transportLayer) Cancel() {
	select {
	case <-tpl.canceled:
		return
	default:
	}

	tpl.cancelOnce.Do(func() {
		close(tpl.canceled)

		tpl.Log().Debug("transport layer canceled")
	})
}

func (tpl *transportLayer) String() string {
	if tpl == nil {
		return "<nil>"
	}

	return fmt.Sprintf("transport.Layer<%s>", tpl.Log().Fields())
}

func (tpl *transportLayer) Done() <-chan struct{} {
	return tpl.done
}

func (tpl *transportLayer) Messages() <-chan sip.Message {
	return tpl.msgs
}

func (tpl *transportLayer) Errors() <-chan error {
	return tpl.errs
}

func (tpl *transportLayer) IsReliable(network string) bool {
	if protocol, ok := tpl.protocols.get(protocolKey(network)); ok && protocol.Reliable() {
		return true
	}
	return false
}

func (tpl *transportLayer) IsStreamed(network string) bool {
	if protocol, ok := tpl.protocols.get(protocolKey(network)); ok && protocol.Streamed() {
		return true
	}
	return false
}

// 日志对象
func (tpl *transportLayer) Log() log.Logger {
	return tpl.log
}
