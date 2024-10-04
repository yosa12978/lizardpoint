package repos

import (
	"context"

	"github.com/google/uuid"
)

type PermissionRepo interface {
	GetChannelPermissions()
	AddRoleToUser(ctx context.Context, userId uuid.UUID, )
}
