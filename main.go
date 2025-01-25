package main

import (
	"GodotServer1/connection"
	"GodotServer1/controller"
	"GodotServer1/parser"
	"log"
	"net/http"
)

func main() {
	ctr := controller.NewController()
	parser := parser.NewMessageParser(ctr)

	connectionHandler := connection.NewConnectionHandler(parser.GetQueueAndStartRead())
	ctr.SetSender(connectionHandler)

	http.HandleFunc("/ws", connectionHandler.HandleConnection)
	// Define the address and port to bind to
	address := ":7070"
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
