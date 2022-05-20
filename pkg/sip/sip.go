package sip

import (
	"fmt"
	"log"
	"net"
)

type SipOption struct {
	address string
	port    int
	udp     bool
	tcp     bool
	tls     bool
}

// 消息回调
type SipCallback func(msg string, info SipInfo)

type SipInfo struct {
	protocol string
	address  string
	port     int
}

var sipOption SipOption

// 初始化
func Start(options SipOption, callback SipCallback) {
	create(options, callback)
}

// 发送消息
func Send() {

}

// 回调监听
func OnRequest(callback SipCallback) {

}

func create(options SipOption, callback SipCallback) {
	makeTransport(options, callback)
}

// 创建通道
func makeTransport(options SipOption, callback SipCallback) {
	if options.udp {
		makeUdpTransport(options, callback)
	}

	if options.tcp {
		makeTcpTransport(options, callback)
	}
}

// udp
func makeUdpTransport(options SipOption, callback SipCallback) (*net.UDPConn, error) {
	address := "0.0.0.0"
	port := 5060
	if options.address != "" {
		address = options.address
	}
	if options.port != 0 {
		port = options.port
	}

	listen, err := net.ListenUDP("upd", &net.UDPAddr{
		IP:   net.IP(address),
		Port: port,
	})
	if err != nil {
		log.Fatalf("Sip Udb Create Failed!")
		return nil, err
	}

	return listen, nil
}

// tcp
func makeTcpTransport(options SipOption, callback SipCallback) (*net.TCPListener, error) {
	address := "0.0.0.0"
	port := 5060
	if options.address != "" {
		address = options.address
	}
	if options.port != 0 {
		port = options.port
	}

	listener, err := net.ListenTCP("upd", &net.TCPAddr{
		IP:   net.IP(address),
		Port: port,
	})

	if err != nil {
		log.Fatalf("Sip Tcp Create Failed!")
		return nil, err
	}

	for {
		// 建立socket连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listen.Accept failed,err:", err)
			continue
		}

		// 业务处理逻辑
		go processTcp(conn, callback)
	}
}

func processTcp(conn net.Conn, callback SipCallback) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("Read from tcp server failed,err:", err)
			break
		}
		data := string(buf[:n])
		fmt.Printf("Recived from client,data:%s\n", data)
	}
}
