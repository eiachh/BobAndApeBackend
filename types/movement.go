package types

import (
	"encoding/json"
	"log"
)

type MoveCommandType string
type AreaName string

const (
	Move     MoveCommandType = "move"
	Teleport MoveCommandType = "teleport"
	Sync     MoveCommandType = "sync"

	LobbyArea           AreaName = "lobby"
	GorillaWarfareArena AreaName = "gorillawarfarearena"
)

const (
	MoveCommandName      CommandName = "MoveCommand"
	AreaEnterCommandName CommandName = "LoginCommand"
)

type MoveCommand struct {
	MoveCommandType MoveCommandType `json:"movecommandtype"`
	PosX            int             `json:"posx"`
	PosY            int             `json:"posy"`
}

type AreaEnterCommand struct {
	AreaName string `json:"areaname"`
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

func NewAreaEnterCmd(body any) *AreaEnterCommand {
	bodyStr, _ := json.Marshal(body)

	var areaEntercmd AreaEnterCommand
	err := json.Unmarshal(bodyStr, &areaEntercmd)

	if err != nil {
		log.Println(err)
	}

	return &areaEntercmd
}
