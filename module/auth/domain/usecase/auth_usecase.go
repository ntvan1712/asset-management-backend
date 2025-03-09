package usecase

import (
	"asset_management_backend/module/auth/domain/entity"
	"context"
)

type AuthUsecase interface {
	Login(context context.Context, loginRequest entity.LoginRequest) (*entity.LoginSuccessResponseEntity, error)
}
