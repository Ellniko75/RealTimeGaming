package main

import (
	"bytes"
	"image/jpeg"
	"log"
	"time"

	"github.com/nfnt/resize"
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
func sendCompressedScreenshotToChannel(ch chan bytes.Buffer) {
	for {
		var compressed bytes.Buffer
		img, _ := screenshot.CaptureScreen()

		//timeSt := time.Now()

		resized := resize.Resize(1280, 720, img, resize.NearestNeighbor)
		errJPG := jpeg.Encode(&compressed, resized, nil)
		if errJPG != nil {
			log.Println("error on encoding jpg")
		}
		//elapsed := time.Since(timeSt)
		//fmt.Println(elapsed)

		ch <- compressed

		time.Sleep(1 * time.Millisecond)
	}
}
