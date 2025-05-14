package datasource

import (
	"asset_management_backend/module/borrowed_asset/data/model"
	"context"
)

type BorrowedAssetDataSource interface {
	ReturnAsset(ctx context.Context, borrowedAssetId int) (*model.BorrowedAsset, error)
}
