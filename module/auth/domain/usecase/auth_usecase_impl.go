package usecase

import (
	"asset_management_backend/common/enums"
	"asset_management_backend/module/auth/repository"
	"context"
)

type authUsecase struct {
	authRepository repository.AuthRepository
}

// HasUserAuthority implements AuthUsecase.
func (a *authUsecase) HasUserAuthority(context context.Context, userID int, userAuthority string) (bool, error) {
	if userAuthority == enums.UserAuthority.Employee {
		return true, nil
	}

	user, err := a.authRepository.FindUserAuthByID(context, userID)
	if err != nil {
		return false, err
	}

	return user.HasAuthority(userAuthority), nil
}

func NewAuthUsecase() AuthUsecase {
	return &authUsecase{
		authRepository: repository.NewAuthRepository(),
	}
}
