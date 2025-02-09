package parser

import (
	"GodotServer1/controller"
	"GodotServer1/types"
	"encoding/json"
	"log"

	"github.com/google/uuid"
)

type MessageParser struct {
	controller     *controller.Controller
	parsingQueue   chan (types.PackageWithUuid)
	isReadingQueue bool
}

func NewMessageParser(ctr *controller.Controller) *MessageParser {
	return &MessageParser{
		controller:   ctr,
		parsingQueue: make(chan types.PackageWithUuid),
	}
}

func (msgParser *MessageParser) GetQueueAndStartRead() chan<- types.PackageWithUuid {
	if !msgParser.isReadingQueue {
		go func() {
			for uuidPkg := range msgParser.parsingQueue {
				msgParser.ParseIncomingMsg(uuidPkg) // Handle messages
			}
		}()
	}
	return msgParser.parsingQueue
}

func (msgParser *MessageParser) ParseIncomingMsg(uuidPkg types.PackageWithUuid) {
	log.Printf("Parsing msg: %s", uuidPkg.BytePackage)
	incomingPkg, err := BytePkgToPkg(uuidPkg)

	if err != nil {
		log.Println(err)
	}

	if incomingPkg.Name == types.LoginCommandName {
		cmd := types.NewLoginCmd(incomingPkg.Body)
		if cmd.UserId == uuid.Nil {
			return
		}
		msgParser.controller.AddAsLoggedIn(cmd)

	} else if incomingPkg.Name == types.MoveCommandName {
		cmd := types.NewMoveCmd(incomingPkg.Body)
		msgParser.controller.Move(uuidPkg.UserId, cmd)
	} else if incomingPkg.Name == types.AreaEnterCommandName {
		cmd := types.NewAreaEnterCmd(incomingPkg.Body)
		msgParser.controller.AreaEnter(uuidPkg.UserId, cmd)
	}
}

func BytePkgToPkg(msg types.PackageWithUuid) (types.Package, error) {
	var incomingPkg types.Package
	err := json.Unmarshal(msg.BytePackage, &incomingPkg)
	if err != nil {
		return incomingPkg, err
	}
	return incomingPkg, nil
}
