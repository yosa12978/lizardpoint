package types

import (
	"github.com/google/uuid"
)

type Channel struct {
	Id               uuid.UUID
	Name             string
	ReadPermissions  []Role
	WritePermissions []Role
}
