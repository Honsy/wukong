package rtsp

import (
	"fmt"
	"test/lib"
)

type CurrentAuthenticator interface {
	realm() string
	username() string
	password() string
	nonce() string
	computeDigestResponse(string, string) string
}

type currentAuthenticator struct {
	fRealm    string
	fUsername string
	fPassword string
	fNonce    string
}

func (ca currentAuthenticator) realm() string {
	return ca.fRealm
}

func (ca currentAuthenticator) username() string {
	return ca.fUsername
}

func (ca currentAuthenticator) password() string {
	return ca.fPassword
}
func (ca currentAuthenticator) nonce() string {
	return ca.fNonce
}

func (ca currentAuthenticator) computeDigestResponse(cmd string, url string) string {
	// The "response" field is computed as:
	//    md5(md5(<username>:<realm>:<password>):<nonce>:md5(<cmd>:<url>))
	// or, if "fPasswordIsMD5" is True:
	//    md5(<password>:<nonce>:md5(<cmd>:<url>))
	HA1 := fmt.Sprintf("%s:%s:%s", ca.fUsername, ca.fRealm, ca.fPassword)
	HA2 := fmt.Sprintf("%s:%s:%s", lib.MD5(HA1), ca.fNonce, lib.MD5(fmt.Sprintf("%s:%s", cmd, url)))
	HA3 := lib.MD5(HA2)
	return HA3
}
