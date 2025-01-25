package types

import (
	"encoding/json"
	"log"
)

type MoveCommandType string

const (
	Move     MoveCommandType = "move"
	Teleport MoveCommandType = "teleport"
	Sync     MoveCommandType = "sync"
)

const (
	MoveCommandName CommandName = "MoveCommand"
)

type MoveCommand struct {
	MoveCommandType MoveCommandType `json:"movecommandtype"`
	PosX            int             `json:"posx"`
	PosY            int             `json:"posy"`
}

func NewMoveCmd(body any) *MoveCommand {
	bodyStr, _ := json.Marshal(body)

	var moveCmd MoveCommand
	err := json.Unmarshal(bodyStr, &moveCmd)

	if err != nil {
		log.Println(err)
	}

	return &moveCmd
}
