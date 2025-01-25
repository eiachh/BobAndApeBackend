package types

import "github.com/google/uuid"

type CommandName string

type Package struct {
	Name CommandName `json:"name"`
	Body any         `json:"body"`
}

type PackageWithUuid struct {
	UserId      uuid.UUID `json:"-"`
	BytePackage []byte    `json:"-"`
}
