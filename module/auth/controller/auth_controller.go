package controller

import (
	"asset_management_backend/module/auth/domain/usecase"
)

type AuthController struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthController() *AuthController {
	return &AuthController{
		authUsecase: usecase.NewAuthUsecase(),
	}
}
