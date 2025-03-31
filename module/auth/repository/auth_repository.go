package repository

import (
	"asset_management_backend/module/auth/domain/entity"
	"context"
)

type AuthRepository interface {
	FindUserAuthByID(ctx context.Context, userID int) (*entity.UserAuthEntity, error)
}
