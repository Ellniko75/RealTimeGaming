package main

import (
	"bytes"
	"image/jpeg"
	"log"
	"sort"
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
		img, err := screenshot.CaptureScreen()
		//timeSt := time.Now()
		resized := resize.Resize(800, 600, img, resize.NearestNeighbor)
		err = jpeg.Encode(&compressed, resized, nil)
		//elapsed := time.Since(timeSt)
		//sfmt.Println(elapsed)

		if err != nil {
			log.Print("ERROOOOOOOOR")
		}
		ch <- compressed

		time.Sleep(1 * time.Millisecond)
	}
}

// returns the index of the item, if not exists returns -1
func contains(slice []repeatedItems, item byte) int {
	for i, element := range slice {
		if element.number == item {
			return i
		}
	}
	return -1
}

type repeatedItems struct {
	number      byte
	times       uint32
	huffmanCode string
}

type tree struct {
	nodeValue    byte
	repeatedTime byte
	left         *tree
	right        *tree
}

func createTree(arr []repeatedItems) {

	for i := 2; i < len(arr); i++ {

	}

}

func sendCompressedHuffman() {
	numbersRepeated := []repeatedItems{}
	for {
		img, _ := screenshot.CaptureScreen()

		for i := 0; i < len(img.Pix); i++ {
			index := contains(numbersRepeated, img.Pix[i])
			//if not exists in slice
			if index == -1 {
				newItem := repeatedItems{number: img.Pix[i], times: 1}
				numbersRepeated = append(numbersRepeated, newItem)
			} else {
				numbersRepeated[index].times++
			}
		}
		//sort increasingly
		sort.Slice(numbersRepeated, func(i, j int) bool {
			return numbersRepeated[i].times < numbersRepeated[j].times
		})
		//create huffman tree
		createTree(numbersRepeated)
		numbersRepeated = []repeatedItems{}

		return
	}

}
