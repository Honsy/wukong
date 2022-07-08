package rtsp

import (
	"fmt"
	"net"
	"test/pkg/logging"
)

var (
	fBaseURL                    string               // 根Url
	fRtspOption                 RTSPOptions          // rtsp配置
	fRtspClient                 RTSPClient           // rtsp对象
	fConnection                 net.Conn             // Socket Connection
	fRequestsAwaitingConnection []RequestRecord      // 存放等待请求的数组
	connectionIsPending         bool                 // 请求Pending标志位
	fInputSocketNum             int                  // 是否成功发起Socket连接 0 1 -1
	fCurrentAuthenticator       CurrentAuthenticator // RTSP认证
	fUserAgentHeader            string               // UA Header
)

type RTSPClient interface {
	sendDescribeCommand()
	ContinueAfterClientCreation1()
}

type rtspClient struct {
}

func Setup(rtspURL string) {
	fRtspClient = NewRtspClient(rtspURL)
}

// 初始化实例
func NewRtspClient(rtspURL string) (r RTSPClient) {
	rtspClient := &rtspClient{}
	fRtspOption = ParseRTSPUrl(rtspURL)
	setUserAgentString("GO RTSP")
	return rtspClient
}

func (rc rtspClient) ContinueAfterClientCreation1() {
	rc.sendOptionsCommand()
}

// 设置根Url
func setBaseUrl(rtspURL string) {
	fBaseURL = rtspURL
}

// 打开Socket
func openConnection() error {
	var err error
	fConnection, err = net.Dial("tcp", fRtspOption.address)
	if err != nil {
		logging.Error("RTSP连接失败！", fBaseURL)
		return err
	}

	return nil
}

// 发送Describe命令
func (r rtspClient) sendDescribeCommand() {
	r.sendRequest(NewRequestRecord())
}

// 发送Options命令
func (r rtspClient) sendOptionsCommand() {
	r.sendRequest(NewRequestRecord())
}

// 发送Announce命令
func (r rtspClient) sendAnnounceCommand() {
	r.sendRequest(NewRequestRecord())
}

// 发送Setup命令
func (r rtspClient) sendSetupCommand() {
	r.sendRequest(NewRequestRecord())
}

// 发送Play命令
func (r rtspClient) sendPlayCommand() {
	r.sendRequest(NewRequestRecord())
}

// 发送Pause命令
func (r rtspClient) sendPauseCommand() {
	r.sendRequest(NewRequestRecord())
}

// 发送Record命令
func (r rtspClient) sendRecordCommand() {
	r.sendRequest(NewRequestRecord())
}

// 发送Teardown命令
func (r rtspClient) sendTeardownCommand() {
	r.sendRequest(NewRequestRecord())
}

// 发送SetParameter命令
func (r rtspClient) sendSetParameterCommand() {
	r.sendRequest(NewRequestRecord())
}

// 发送GetParameter命令
func (r rtspClient) sendGetParameterCommand() {
	r.sendRequest(NewRequestRecord())
}

func (r rtspClient) sendRequest(request RequestRecord) string {
	if len(fRequestsAwaitingConnection) != 0 {
		connectionIsPending = true
	} else if fInputSocketNum < 0 {
		err := openConnection()
		if err != nil {
			return ""
		}
	}

	if connectionIsPending {
		fRequestsAwaitingConnection = append(fRequestsAwaitingConnection, request)
		return request.cseq()
	}

	var cmdURL = fBaseURL          // by default
	const protocolStr = "RTSP/1.0" // by default
	var contentLengthHeaderFmt = ""
	var extraHeaders = SetRequsetFields(request, cmdURL, protocolStr, "")
	var contentStr = request.contentStr()
	if contentStr != "" {
		contentLengthHeaderFmt = fmt.Sprintf("Content-Length: %d\r\n", len(contentStr))
	}
	var authenticatorStr = createAuthenticatorString(request.commandName(), fBaseURL)

	var cmdFmt = fmt.Sprintf("%s %s %s\r\n", request.commandName(), fBaseURL, protocolStr) +
		fmt.Sprintf("CSeq: %d\r\n", request.cseq()) +
		authenticatorStr +
		fUserAgentHeader +
		extraHeaders +
		contentLengthHeaderFmt +
		"\r\n" +
		contentStr

	_, err := fConnection.Write([]byte(cmdFmt))

	if err != nil {
		return err.Error()
	}
	return ""
}
