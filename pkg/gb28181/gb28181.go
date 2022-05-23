package gb28181

import (
	"test/log"
	"test/pkg/sip"
	"test/pkg/sip/sipserver"
)

var (
	logger log.Logger
	server sipserver.Server
)

func init() {
	logger = log.NewDefaultLogrusLogger().WithPrefix("SipServer")
	// logger.SetLevel(log.TraceLevel)
	logger.SetLevel(log.WarnLevel)
}

// 启动sip服务器
func Setup() {
	server = sipserver.NewServer(sipserver.ServerConfig{}, nil, nil, logger)
	server.Listen("udp", "0.0.0.0:5061", nil)
	server.OnRequest("INVITE", INVITE)
	server.OnRequest("REGISTER", REGISTER)

}

func INVITE(req sip.Request, tx sip.ServerTransaction) {
}

// 设备注册
func REGISTER(req sip.Request, tx sip.ServerTransaction) {
	header := make([]sip.Header, 0)
	auth := sip.AuthFromValue(req.GetHeaders("Authorization")[0].Value())
	logger.Debugf("_____________DDDDDDDD", auth.Username())

	// 判断请求头是否哦包含设备认证消息
	if auth.Username() != "" {
		server.RespondOnRequest(req, 401, "sss", "", header)

		return
	}

	server.RespondOnRequest(req, 401, "sss", "", header)
}
