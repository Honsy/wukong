package transport

import (
	"fmt"
	"strings"
	"test/log"
	"test/pkg/sip"
	"time"
)

const (
	//netErrRetryTime = 5 * time.Second
	sockTTL = time.Hour
)

type ProtocolFactory func(
	network string,
	output chan<- sip.Message,
	errs chan<- error,
	cancel <-chan struct{},
	msgMapper sip.MessageMapper,
	log log.Logger,
) (Protocol, error)

type protocol struct {
	network  string
	reliable bool
	streamed bool
	log      log.Logger
}

type Protocol interface {
	Done() <-chan struct{}
	Network() string
	Reliable() bool
	Streamed() bool
	Listen(targer *Target, options ...ListenOption) error
	Send(target *Target, msg sip.Message) error
	String() string
}

func (pr *protocol) Log() log.Logger {
	return pr.log
}

func (pr *protocol) String() string {
	if pr == nil {
		return "<nil>"
	}

	fields := pr.Log().Fields().WithFields(log.Fields{
		"network": pr.network,
	})

	return fmt.Sprintf("transport.Protocol<%s>", fields)
}

// Common Network
func (pr *protocol) Network() string {
	return strings.ToUpper(pr.network)
}

func (pr *protocol) Reliable() bool {
	return pr.reliable
}

func (pr *protocol) Streamed() bool {
	return pr.streamed
}
