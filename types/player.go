package types

import (
	"github.com/google/uuid"
)

type Player struct {
	UserID uuid.UUID `json:"userid"`
	Name   string    `json:"name"`
}
