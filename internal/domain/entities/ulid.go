package entities

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

func NewULID() uuid.UUID {
	bytes := ulid.Make().Bytes()
	id, _ := uuid.FromBytes(bytes)
	return id
}
