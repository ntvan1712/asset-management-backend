package usecase

import (
	"asset_management_backend/module/auth/domain/entity"
	"asset_management_backend/module/auth/repository"
	"context"
)

type authUsecase struct {
	authRepository repository.AuthRepository
}

// Login implements AuthUsecase.
func (a *authUsecase) Login(context context.Context, loginRequest entity.LoginRequest) (*entity.LoginSuccessResponseEntity, error) {
	return a.authRepository.Login(context, loginRequest.UserName, loginRequest.Password)

}

func NewAuthUsecase() AuthUsecase {
	return &authUsecase{
		authRepository: repository.NewAuthRepository(),
	}
}
