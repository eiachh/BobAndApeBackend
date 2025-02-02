package connection

import (
	"GodotServer1/types"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ConnectionHandler struct {
	IncommingMsg chan<- types.PackageWithUuid
	Connections  map[uuid.UUID]*websocket.Conn
}

func NewConnectionHandler(incommingMsgs chan<- types.PackageWithUuid) *ConnectionHandler {

	return &ConnectionHandler{IncommingMsg: incommingMsgs}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (can be restricted for security)
	},
}

func (conHandler *ConnectionHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()
	log.Println("New WebSocket connection established.")
	log.Println("Reading username request")
	var userID uuid.UUID
	for {
		uID, loginErr := conHandler.HandleLogin(conn)
		if loginErr == nil {
			userID = uID
			break // Exit the loop when login is successful
		}
	}

	for {
		connType, msg, err := conn.ReadMessage()
		if err != nil {

			conHandler.RemoveConnection(conn)
			log.Println("Error reading message:", err)
			break
		}

		if connType != 1 {
			log.Println("The incoming msg was not text based")
			break
		}
		conHandler.IncommingMsg <- types.PackageWithUuid{UserId: userID, BytePackage: msg}
	}
}

func (conHandler *ConnectionHandler) RegisterPlayerConn(conn *websocket.Conn, u uuid.UUID) {
	if conHandler.Connections == nil {
		conHandler.Connections = make(map[uuid.UUID]*websocket.Conn)
	}
	conHandler.Connections[u] = conn
}

func (conHandler *ConnectionHandler) RemoveConnection(connToRemove *websocket.Conn) {
	for uid, conn := range conHandler.Connections {
		if conn == connToRemove {
			delete(conHandler.Connections, uid)
		}
	}
}

func (conHandler *ConnectionHandler) HandleLogin(conn *websocket.Conn) (uuid.UUID, error) {
	connType, msg, err := conn.ReadMessage()
	if connType != 1 {
		return uuid.Nil, errors.New("not plain text was used for login")
	}
	if err != nil {
		return uuid.Nil, err
	}

	var incomingPkg types.Package
	jsonerr := json.Unmarshal(msg, &incomingPkg)
	if jsonerr != nil {
		return uuid.Nil, err
	}
	if incomingPkg.Name != types.LoginCommandName {
		return uuid.Nil, errors.New("this was not a login message")
	}
	loginCmd := types.NewLoginCmd(incomingPkg.Body)
	loginCmd.UserId = uuid.New()

	pkgBytes, _ := json.Marshal(types.Package{Name: types.LoginCommandName, Body: loginCmd})
	conHandler.IncommingMsg <- types.PackageWithUuid{UserId: loginCmd.UserId, BytePackage: pkgBytes}

	conHandler.sendPkgTo(conn, types.Package{
		Name: types.LoginAcceptCommandName,
		Body: types.LoginAccept{ReceivedUserID: loginCmd.UserId}})
	return loginCmd.UserId, nil

}
