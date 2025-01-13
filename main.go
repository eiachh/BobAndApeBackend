package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Define a WebSocket upgrader to handle upgrading HTTP connections to WebSocket.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (can be restricted for security)
	},
}

func twoSum(nums []int, target int) []int {
	for i, elem := range nums {
		for j := i + 1; i < len(nums)-1; j++ {
			if elem+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {

	twoSum([]int{3, 2, 3}, 6)

	http.HandleFunc("/ws", handleConnection)
	// Define the address and port to bind to
	address := "127.0.0.1:6969"
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	// Listen on the specified address and port
}

// Handler to handle WebSocket connections
func handleConnection(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	log.Println("New WebSocket connection established.")

	for {
		// Read a message from the WebSocket connection
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Print the received message
		fmt.Println("Received:", string(p))

		// Optionally, send a response back to the client
		if err := conn.WriteMessage(messageType, []byte("Message received!")); err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}
