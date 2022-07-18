package rtsp

import (
	"bufio"
	"fmt"
	"strconv"
	"test/pkg/logging"
)

type StatusCode int

type HeaderValue []string

// Header
type Header map[string]HeaderValue

// RTSP Response
type Response struct {
	// 状态码
	StatusCode StatusCode

	// 消息
	StatusMessage string

	// 请求头
	Header Header

	// 请求体
	Body []byte
}

func (res *Response) Read(rb *bufio.Reader) error {
	byts, err := readBytesLimited(rb, ' ', 255)
	if err != nil {
		return err
	}
	// proto := byts[:len(byts)-1]
	byts, err = readBytesLimited(rb, ' ', 4)
	if err != nil {
		return err
	}
	statusCodeStr := string(byts[:len(byts)-1])

	statusCode64, err := strconv.ParseInt(statusCodeStr, 10, 32)
	if err != nil {
		return fmt.Errorf("unable to parse status code")
	}
	logging.Debug("dd2d2")
	res.StatusCode = StatusCode(statusCode64)
	return nil
}
