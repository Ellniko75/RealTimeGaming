package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Channel that holds all the screenshots
var screenshotsChannel chan []uint8 = make(chan []uint8, 24)

// upgrader to change http to websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// upgrades the http connection to a websocket, calls handleWebSocket to handle the websocket connection
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Print("CLIENT CONNECTED")
	handleWebSocket(ws)
}

// handles the websocket connection
func handleWebSocket(conn *websocket.Conn) {
	errcount := 0
	for {
		//get img from channel
		img := <-screenshotsChannel
		err := conn.WriteMessage(2, img)
		if err != nil {
			log.Println("Posible disconnection: error at handleWebSocket()", err)
			errcount = errcount + 1
		}
		if errcount > 30 {
			return
		}
		time.Sleep(16 * time.Millisecond)
	}
}

func main() {
	go sendScreenshotToChannel(screenshotsChannel)
	http.HandleFunc("/", wsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
