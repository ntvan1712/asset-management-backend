package datasource

import (
	"asset_management_backend/common/error_app"
	"asset_management_backend/infras"
	"asset_management_backend/module/auth/data/model"
	"context"
	"database/sql"
	"errors"
)

type authDataSourceImpl struct {
	dbProvider *infras.DbProvider
}

// FindUserAuthByID implements AuthDataSource.
func (a *authDataSourceImpl) FindUserAuthByID(ctx context.Context, userID int) (*model.UserAuthModel, error) {
	var userAuth model.UserAuthModel

	err := a.dbProvider.Instance.NewSelect().
		Table("users").
		ColumnExpr("users.id, users.role_id").
		ColumnExpr("COALESCE(array_agg(up.permission_id) FILTER (WHERE up.permission_id IS NOT NULL), '{}') AS permission_ids").
		Join("LEFT JOIN user_permissions AS up ON up.user_id = users.id").
		Where("users.id = ?", userID).
		Group("users.id", "users.role_id").
		Scan(ctx, &userAuth)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, error_app.ErrDocumentNotFound
		}
		return nil, err
	}

	return &userAuth, nil

}

func NewAuthDataSource(dbProvider *infras.DbProvider) AuthDataSource {
	return &authDataSourceImpl{
		dbProvider: dbProvider,
	}
}
