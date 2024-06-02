package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Channel that holds the screenshots taken (40 at max)
var screenshotsChannel = make(chan bytes.Buffer, 40)

// upgrader to change http to websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// FIRST ENTRY POINT: upgrades the http connection to a websocket, calls handleWebSocket to handle the websocket connection
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Print("CLIENT CONNECTED")
	go handleWebSocket(ws)
}

// handles the websocket connection
func handleWebSocket(conn *websocket.Conn) {
	go manageClientInputs(conn)
	errcount := 0
	for {
		//get img from channel and send it to the client via websocket
		img := <-screenshotsChannel
		err := conn.WriteMessage(2, img.Bytes())

		if err != nil {
			errcount = errcount + 1
		}
		if errcount > 30 {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
}

type jsonKeyInputs struct {
	Key    string
	Action string
}

// reads messages from the client on a loop and executes them calling the presskey function
func manageClientInputs(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var deseralizedJson jsonKeyInputs
		errorr := json.Unmarshal(msg, &deseralizedJson)
		if errorr != nil {
			fmt.Print("errorrrrrrrr", err)
		}
		if deseralizedJson.Action == "down" {
			//key down
			keyDownRobotgo(deseralizedJson.Key)
		} else {
			//key up
			keyUpRobotgo(deseralizedJson.Key)
		}
		//ExecuteKey(string(msg))
		//log.Print("client message: ", string(msg))
		time.Sleep(20 * time.Millisecond)
	}

}

func main() {
	//go sendCompressedScreenshotToChannel(screenshotsChannel)
	go sendCompressedScreenshotToChannel(screenshotsChannel)
	//sendCompressedHuffman()
	http.HandleFunc("/", wsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
