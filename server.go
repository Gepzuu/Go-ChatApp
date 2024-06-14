package main

import (
	"log"

	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{

		CheckOrigin: func(r *http.Request) bool {

			return true

		},
	}
)

func main() {

	http.HandleFunc("/ws", handleWebSocket)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {

		log.Println("Error upgrading to WebSocket:", err)

		return

	}

	defer conn.Close()

	for {

		messageType, msg, err := conn.ReadMessage()

		if err != nil {

			log.Println("Error reading message:", err)

			break

		}

		if err := conn.WriteMessage(messageType, msg); err != nil {

			log.Println("Error writing message:", err)

			break

		}

	}

}
