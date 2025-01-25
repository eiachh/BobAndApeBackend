package controller

import (
	"GodotServer1/types"

	"github.com/google/uuid"
)

type Sender interface {
	SendTo(to uuid.UUID, pkg types.Package) error
	SendAll(pkg types.Package) error
	SendExcept(except uuid.UUID, pkg types.Package) error
}

type Controller struct {
	players map[uuid.UUID]types.Player
	sender  Sender
}

func NewController() *Controller {
	return &Controller{
		players: make(map[uuid.UUID]types.Player),
	}
}

func (controller *Controller) SetSender(newSender Sender) {
	controller.sender = newSender
}

func (controller *Controller) Move(userId uuid.UUID, cmd *types.MoveCommand) {

}

func (controller *Controller) AddAsLoggedIn(cmd *types.LoginCommand) {
	controller.players[cmd.UserId] = types.Player{
		UserID: cmd.UserId,
		Name:   cmd.NameRequest,
	}
}
