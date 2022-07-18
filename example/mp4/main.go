package main

import (
	"os"
	"os/signal"
	"time"
)

func main() {

	// mp4Client := mp4.NewMp4Client("")

	c := make(chan os.Signal)
	signal.Notify(c)
	go func() {
		time.Sleep(3 * time.Second)
	}()
	<-c
}
