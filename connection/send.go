package connection

import (
	"GodotServer1/types"
	"encoding/json"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (conHandler *ConnectionHandler) SendTo(to uuid.UUID, pkg types.Package) error {
	conn, hasKey := conHandler.Connections[to]
	if !hasKey {
		return errors.New("no key for uuid found in connections")
	}
	conHandler.sendPkgTo(conn, pkg)
	return nil
}
func (conHandler *ConnectionHandler) SendAll(pkg types.Package) error {
	if len(conHandler.Connections) <= 0 {
		return errors.New("there is no client connected")
	}
	for _, conn := range conHandler.Connections {
		conHandler.sendPkgTo(conn, pkg)
	}
	return nil
}
func (conHandler *ConnectionHandler) SendExcept(except uuid.UUID, pkg types.Package) error {
	return nil
}

func (conHandler *ConnectionHandler) sendPkgTo(conn *websocket.Conn, pkg types.Package) {
	log.Println("Sending back response")
	pkgStr, _ := json.Marshal(pkg)
	log.Println(pkgStr)
	// Type 1 is txt message
	if err := conn.WriteMessage(1, pkgStr); err != nil {
		log.Println("Error sending message:", err)
	}
}
