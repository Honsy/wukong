package main

import (
	"os"
	"os/signal"
	"test/pkg/rtsp"
	"time"
)

func main() {
	rtspClient := rtsp.NewRtspClient("rtsp://admin:aykj1234@10.0.16.190:554/Streaming/Channels/101?transportmode=unicast")

	rtspClient.ContinueAfterClientCreation1()

	c := make(chan os.Signal)
	signal.Notify(c)
	go func() {
		time.Sleep(3 * time.Second)
	}()
	<-c
}
