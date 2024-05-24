package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vova616/screenshot"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Print("CLIENT CONNECTED")
	handleWebSocket(ws)
}
func handleWebSocket(conn *websocket.Conn) {
	errcount := 0
	for {
		img, _ := screenshot.CaptureScreen()

		err := conn.WriteMessage(2, img.Pix)
		if err != nil {
			log.Println("Posible disconnection: error at handleWebSocket()", err)
			errcount = errcount + 1
		}
		if errcount > 100 {
			return
		}
		time.Sleep(16 * time.Millisecond)
	}
}

func main() {
	http.HandleFunc("/", wsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
