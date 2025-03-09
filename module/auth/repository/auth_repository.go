package repository

import (
	"asset_management_backend/module/auth/domain/entity"
	"context"
)

type AuthRepository interface {
	Login(context context.Context, username string, password string) (*entity.LoginSuccessResponseEntity, error)
}
