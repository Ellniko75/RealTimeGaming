package main

import (
	"log"
	"time"

	"github.com/vova616/screenshot"
)

func sendScreenshotToChannel(ch chan []uint8) {
	for {
		img, err := screenshot.CaptureScreen()
		if err != nil {
			log.Print("ERROOOOOOOOR")
		}
		if len(ch) == 24 {
			<-ch
		}
		ch <- img.Pix
		time.Sleep(16 * time.Millisecond)
	}

}
