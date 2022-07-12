package main

import (
	"fmt"
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
		fmt.Println("Go routine running")
		time.Sleep(3 * time.Second)
		fmt.Println("Go routine done")
	}()
	<-c
	fmt.Println("bye")
}
