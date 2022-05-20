package gb28181

import (
	"test/log"
	"test/pkg/sip"
	"test/pkg/sip/sipserver"
)

var (
	logger log.Logger
)

func init() {
	logger = log.NewDefaultLogrusLogger().WithPrefix("SipServer")
	// logger.SetLevel(log.TraceLevel)
	logger.SetLevel(log.DebugLevel)
}

// 启动sip服务器
func Setup() {
	server := sipserver.NewServer(sipserver.ServerConfig{}, nil, nil, logger)
	server.Listen("udp", "0.0.0.0:5081", nil)
	server.OnRequest("INVITE", INVITE)
	server.OnRequest("REGISTER", REGISTER)

}

func INVITE(req sip.Request, tx sip.ServerTransaction) {
	logger.Debug("DDDDDDDDDDDDDD")
}

// 设备注册
func REGISTER(req sip.Request, tx sip.ServerTransaction) {
	logger.Debug("DDDDDDDDDDDDDD")
}
