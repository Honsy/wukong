package rtsp

import (
	"fmt"
	"strings"
	"test/lib"
)

type RequestRecord interface {
	cseq() string
	commandName() string
	contentStr() string
}

type requestRecord struct {
	fCSeq        string
	fCommandName string
	fContentStr  string
}

func NewRequestRecord() (r RequestRecord) {
	rr := &requestRecord{}
	return rr
}

// 设置请求字段
func SetRequsetFields(r RequestRecord, cmdUrl string, protocolStr string, extraHeaders string) string {
	if strings.ToUpper(r.commandName()) == "DESCRIBE" {
		extraHeaders = "Accept: application/sdp\r\n"
	} else if strings.ToUpper(r.commandName()) == "OPTIONS" {
		extraHeaders = createSessionString("ddd")
	} else if strings.ToUpper(r.commandName()) == "ANNOUNCE" {
		extraHeaders = "Content-Type: application/sdp\r\n"
	} else if strings.ToUpper(r.commandName()) == "SETUP" {

	}

	return extraHeaders
}

// 创建UA
func setUserAgentString(ua string) {
	if ua == "" {
		return
	}
	fUserAgentHeader = fmt.Sprintf("User-Agent: %s\r\n", ua)
}

// 创建Authenticatior字段
func createAuthenticatorString(cmd string, url string) string {
	var auth = fCurrentAuthenticator
	var authenticatorStr string
	if auth.realm() != "" && auth.username() != "" && auth.password() != "" {
		if auth.nonce() != "" {
			// Digest authentication
			response := auth.computeDigestResponse(cmd, url)
			authenticatorStr = fmt.Sprintf("Authorization: Digest username=%s, realm=%s, nonce=%s, uri=%s, response=%s\r\n", auth.username(), auth.realm(), auth.nonce(), url, response)
		} else {
			// Basic authentication
			response := lib.Base64Encode(fmt.Sprintf("%s:%s", auth.username(), auth.password()))
			authenticatorStr = fmt.Sprintf("Authorization: Basic %s\r\n", response)
		}
	}
	return authenticatorStr
}

func createSessionString(sessionId string) string {
	var sessionStr string
	if sessionId != "" {
		sessionStr = fmt.Sprintf("%d", len(sessionId))
	} else {
		sessionStr = ""
	}
	return sessionStr
}

func (r requestRecord) cseq() string {
	return r.fCSeq
}

func (r requestRecord) commandName() string {
	return r.fCommandName
}

func (r requestRecord) contentStr() string {
	return r.fContentStr
}
