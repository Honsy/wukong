package rtsp

import (
	"bufio"
	"fmt"
	"net"
	"test/pkg/logging"
)

var (
	fBaseURL                    string                                // 根Url
	fRtspOption                 RTSPOptions                           // rtsp配置
	fRtspClient                 RTSPClient                            // rtsp 接口抽象对象
	fRtsp                       rtspClient                            // rtsp 结构体对象
	fConnection                 net.Conn                              // Socket Connection
	fRequestsAwaitingConnection []RequestRecord                       // 存放等待请求的数组
	connectionIsPending         bool                                  // 请求Pending标志位
	fInputSocketNum             int                                   // 是否成功发起Socket连接 0 1 -1
	fCurrentAuthenticator       CurrentAuthenticator                  // RTSP认证
	fUserAgentHeader            string                                // UA Header
	bufferSize                  uint16               = 65535 - 20 - 8 // IPv4 max size - IPv4 Header size - UDP Header size
	fCSeq                       int                                   // 请求序列号
)

type RTSPClient interface {
	sendDescribeCommand()
	ContinueAfterClientCreation1()
	// Requests returns channel with new incoming server transactions.
	Requests() <-chan requestRecord
}

type rtspClient struct {
	requests chan requestRecord
	connRW   *bufio.Reader
}

func Setup(rtspURL string) {
	fRtspClient = NewRtspClient(rtspURL)
}

// 初始化实例
func NewRtspClient(rtspURL string) (r RTSPClient) {
	// 常量赋值
	// fRtsp := &rtspClient{}
	// fRtspOption = ParseRTSPUrl(rtspURL)
	// fCurrentAuthenticator = &currentAuthenticator{
	// 	fUsername: "",
	// 	fPassword: "",
	// 	fRealm:    "",
	// 	fNonce:    "",
	// }
	// fInputSocketNum = -1
	// fCSeq = 0

	// // 开始建立连接
	// fConnection, _ := net.Dial("tcp", fRtspOption.address)

	// timeoutConn := RichConn{
	// 	conn,
	// 	timeout,
	// }

	// bufio.NewReaderSize(fConnection, int(bufferSize))

	// // 设置UA
	// setBaseUrl(rtspURL)
	// setUserAgentString("GO RTSP")
	// return fRtsp
	return nil
}

func (rc rtspClient) ContinueAfterClientCreation1() {
	rc.sendOptionsCommand()
}

func (rc rtspClient) Requests() <-chan requestRecord {
	return rc.requests
}

// 设置根Url
func setBaseUrl(rtspURL string) {
	fBaseURL = rtspURL
}

// 打开Socket
func openConnection() error {
	var err error

	if fRtspOption.username != "" || fRtspOption.password != "" {
		fCurrentAuthenticator.setUsernameAndPassword(fRtspOption.username, fRtspOption.password)
	}

	fConnection, err = net.Dial("tcp", fRtspOption.address)

	if err != nil {
		logging.Error("RTSP连接失败！", fBaseURL)
		fInputSocketNum = -1
		return err
	}

	// timeoutConn := RichConn{
	// 	fConnection,
	// 	time.Second * 6000,
	// }

	// fRtsp.connRW = bufio.NewReadWriter(bufio.NewReaderSize(&timeoutConn, int(bufferSize)), bufio.NewWriterSize(&timeoutConn, int(bufferSize)))

	// logging.Debug("TCP连接成功！正在监听消息...")

	// listenMessages()
	// fInputSocketNum = 1

	return nil
}

// 读取socket信息
func listenMessages() {
	// 循环读取通道内部消息
	// buf := make([]byte, bufferSize)

	go func() {
		for {
			if fConnection != nil {
				// num, err := fConnection.Read(buf)
				// if err != nil {
				// 	logging.Error(err)
				// 	return
				// }
				// logging.Debug("dddd", fConnection)
				// _, err = fRtsp.br.Read(buf[:num])
				// logging.Debug("dddd1", fConnection)
				// if err != nil {
				// 	logging.Error(err)
				// 	return
				// }
				// var res Response
				// logging.Debug("dddd1", fConnection)
				// res.Read(fRtsp.br)

				// data := buf[:num]

				// logging.Debug(string(data))
			}
		}
	}()
}

// 发送Describe命令
func (r rtspClient) sendDescribeCommand() {
	fCSeq += 1
	r.sendRequest(NewRequestRecord(fCSeq, "DESCRIBE", ""))
}

// 发送Options命令
func (r rtspClient) sendOptionsCommand() {
	fCSeq += 1
	r.sendRequest(NewRequestRecord(fCSeq, "OPTIONS", ""))
}

// 发送Announce命令
func (r rtspClient) sendAnnounceCommand() {
	fCSeq += 1
	r.sendRequest(NewRequestRecord(fCSeq, "ANNOUNCE", ""))
}

// 发送Setup命令
func (r rtspClient) sendSetupCommand() {
	fCSeq += 1
	r.sendRequest(NewRequestRecord(fCSeq, "SETUP", ""))
}

// 发送Play命令
func (r rtspClient) sendPlayCommand() {
	fCSeq += 1
	r.sendRequest(NewRequestRecord(fCSeq, "PLAY", ""))
}

// 发送Pause命令
func (r rtspClient) sendPauseCommand() {
	fCSeq += 1
	r.sendRequest(NewRequestRecord(fCSeq, "PAUSE", ""))
}

// 发送Record命令
func (r rtspClient) sendRecordCommand() {
	fCSeq += 1
	r.sendRequest(NewRequestRecord(fCSeq, "RECORD", ""))
}

// 发送Teardown命令
func (r rtspClient) sendTeardownCommand() {
	fCSeq += 1
	r.sendRequest(NewRequestRecord(fCSeq, "TEARDOWN", ""))
}

// 发送SetParameter命令
func (r rtspClient) sendSetParameterCommand() {
	fCSeq += 1
	r.sendRequest(NewRequestRecord(fCSeq, "SETPARAMETER", ""))
}

// 发送GetParameter命令
func (r rtspClient) sendGetParameterCommand() {
	fCSeq += 1
	r.sendRequest(NewRequestRecord(fCSeq, "GETPARAMETER", ""))
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
		fmt.Sprintf("CSeq: %s\r\n", request.cseq()) +
		authenticatorStr +
		fUserAgentHeader +
		extraHeaders +
		contentLengthHeaderFmt +
		"\r\n" +
		contentStr

	_, err := fConnection.Write([]byte(cmdFmt))

	logging.Debug("发送请求：", cmdFmt)
	if err != nil {
		return err.Error()
	}

	// 开始读取请求响应
	// fConnection.Read()

	return ""
}
