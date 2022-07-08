package rtsp

import (
	"strings"
)

type RTSPOptions struct {
	url      string
	username string
	password string
	address  string
	port     int
}

// 解析RTSP URl rtps[s]://[username:password@]ip:port/...
func ParseRTSPUrl(url string) (rOptions RTSPOptions) {
	const rtspPrefix = "rtsp://"
	const rtspPrefixLength = 7
	const rtspsPrefix = "rtsps://"
	const rtspsPrefixLength = 8
	var username string
	var password string
	var address string
	var from string
	// 确定端口
	if strings.HasPrefix(url, rtspPrefix) {
		from = strings.TrimPrefix(url, rtspPrefix)
	} else {
		from = strings.TrimPrefix(url, rtspsPrefix)
	}
	/*
		确定用户名密码 简单点 直接判断url是否包含@符号进行切割
		[username:password@]
	*/
	if strings.Index(from, "@") != -1 {
		splitArray := strings.Split(from, "@")
		from = splitArray[1]
		usernameAndPassword := splitArray[0]
		upArray := strings.Split(usernameAndPassword, ":")
		username = upArray[0]
		password = upArray[1]
	}

	// 直接解析IP 此处 from = ip:port/....
	addressAndSuffixUrlArray := strings.Split(from, "/")
	address = addressAndSuffixUrlArray[0]

	rtspOptions := &RTSPOptions{
		address:  address,
		username: username,
		password: password,
	}

	return *rtspOptions
}
