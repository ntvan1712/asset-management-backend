package repository

import (
	"asset_management_backend/infras"
	datasource "asset_management_backend/module/auth/data/data_source"
	"asset_management_backend/module/auth/domain/entity"
	"context"
)

type authRepository struct {
	authDS datasource.AuthDataSource
}

// FindUserAuthByID implements AuthRepository.
func (a *authRepository) FindUserAuthByID(ctx context.Context, userID int) (*entity.UserAuthEntity, error) {
	authModel, err := a.authDS.FindUserAuthByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return authModel.ToEntity(), nil
}

func NewAuthRepository() AuthRepository {
	return &authRepository{
		authDS: datasource.NewAuthDataSource(infras.GetDbProvider()),
	}
}
