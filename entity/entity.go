package entity

import "github.com/google/uuid"

// ID unique id
type ID = uuid.UUID

// NewID create a new uuid
func NewID() ID {
	return ID(uuid.New())
}
