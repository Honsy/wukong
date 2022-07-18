package mp4

import (
	"log"
	"net/http"
	"time"
)

var (
	bufferSize uint16 = 65535 - 20 - 8 // IPv4 max size - IPv4 Header size - UDP Header size
	buffer     []byte
)

func NewMp4Client(url string) {
	httpClient := &http.Client{Timeout: 5 * time.Second}

	resp, err := httpClient.Get(url)
	if err != nil {
		return
	}
	buffer := make([]byte, bufferSize)

	go readMp4Buffer(resp)
}

func readMp4Buffer(resp *http.Response) {
	for {
		num, err := resp.Body.Read(buffer)

		log.Printf(string(buffer))
	}
}
