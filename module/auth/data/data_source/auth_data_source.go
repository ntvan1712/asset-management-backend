package datasource

import (
	"asset_management_backend/module/auth/data/model"
	"context"
)

type AuthDataSource interface {
	FindUserAuthByID(ctx context.Context, userID int) (*model.UserAuthModel, error)
}
