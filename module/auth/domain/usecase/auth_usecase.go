package usecase

import "context"

type AuthUsecase interface {
	HasUserAuthority(context context.Context, userID int, userAuthority string) (bool, error)
}
