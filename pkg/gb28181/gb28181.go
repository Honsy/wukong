package gb28181

import (
	"fmt"
	"net/http"
	"test/lib"
	"test/log"
	"test/pkg/sip"
	"test/pkg/sip/sipserver"
)

var (
	logger           log.Logger
	server           sipserver.Server
	defaultAlgorithm string
)

func init() {
	defaultAlgorithm = "MD5"
	logger = log.NewDefaultLogrusLogger().WithPrefix("SipServer")
	// logger.SetLevel(log.TraceLevel)
	logger.SetLevel(log.DebugLevel)
}

// 启动sip服务器
func Setup() {
	server = sipserver.NewServer(sipserver.ServerConfig{}, nil, nil, logger)
	server.Listen("udp", "0.0.0.0:5061", nil)
	server.OnRequest(sip.INVITE, INVITE)
	server.OnRequest(sip.REGISTER, REGISTER)
	server.OnRequest(sip.MESSAGE, MESSAGE)
}

func INVITE(req sip.Request, tx sip.ServerTransaction) {

}

//
func MESSAGE(req sip.Request, tx sip.ServerTransaction) {
	body := req.Body()
	message := &MessageReceive{}

	if err := lib.XMLDecode([]byte(body), message); err != nil {
		logger.Errorf("Message Unmarshal xml err:", err, "body:", body)
		tx.Respond(sip.NewResponseFromRequest("", req, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), ""))
		return
	}
	lib.XMLDecode([]byte(req.Body()), message)

	switch message.CmdType {
	case "Keepalive":
	}
}

// 设备注册
func REGISTER(req sip.Request, tx sip.ServerTransaction) {

	// 判断请求头是否哦包含设备认证消息
	if len(req.GetHeaders("Authorization")) == 0 {
		resp := sip.NewResponseFromRequest("", req, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "")
		resp.AppendHeader(&sip.GenericHeader{HeaderName: "WWW-Authenticate", Contents: fmt.Sprintf("Digest Nonce=\"%s\", algorithm=MD5, Realm=\"%s\",qop=\"auth\"", req.MessageID(), req.MessageID())})
		tx.Respond(resp)
		return
	}
	auth := sip.AuthFromValue(req.GetHeaders("Authorization")[0].Value())

	// 判断设备认证消息是否包含User
	if auth.Username() == "" {
		resp := sip.NewResponseFromRequest("", req, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "")
		resp.AppendHeader(&sip.GenericHeader{HeaderName: "WWW-Authenticate", Contents: fmt.Sprintf("Digest Nonce=\"%s\", algorithm=MD5, Realm=\"%s\",qop=\"auth\"", req.MessageID(), req.MessageID())})
		tx.Respond(resp)
		return
	}

	// 响应注册成功
	resp := sip.NewResponseFromRequest("", req, http.StatusOK, http.StatusText(http.StatusOK), "")
	tx.Respond(resp)
}
