package types

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
)

const (
	LoginCommandName CommandName = "LoginCommand"
)

type LoginCommand struct {
	NameRequest string    `json:"namerequest"`
	UserId      uuid.UUID `json:"userid"`
}

func NewLoginCmd(body any) *LoginCommand {
	bodyStr, _ := json.Marshal(body)

	var loginCmd LoginCommand
	err := json.Unmarshal(bodyStr, &loginCmd)

	if err != nil {
		log.Println(err)
	}
	fmt.Println(loginCmd.NameRequest)
	return &loginCmd
}

type LoginAccept struct {
	ReceivedUserID uuid.UUID `json:"receiveduserid"`
}

type LoginRefused struct {
	Message string `json:"message"`
}
