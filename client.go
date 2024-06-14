package main

import (
	"bufio"

	"fmt"

	"log"

	"os"

	"time"

	"github.com/gorilla/websocket"
)

func main() {

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)

	if err != nil {

		log.Fatal("Error connecting to WebSocket:", err)

	}

	defer conn.Close()

	go func() {

		for {

			_, msg, err := conn.ReadMessage()

			if err != nil {

				log.Println("Error reading message:", err)

				return

			}

			fmt.Println("Received:", string(msg))

		}

	}()

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("Enter message: ")

		message, _ := reader.ReadString('\n')

		message = message[:len(message)-1]

		err := conn.WriteMessage(websocket.TextMessage, []byte(message))

		if err != nil {

			log.Println("Error writing message:", err)

			return

		}

		time.Sleep(time.Second) // Delay to allow reading messages

	}

}
