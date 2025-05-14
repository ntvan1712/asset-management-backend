package repository

import (
	"asset_management_backend/module/borrowed_asset/domain/entity"
	"context"
)

type BorrowedAssetRepository interface {
	ReturnAsset(ctx context.Context, borrowedAssetId int) (*entity.BorrowedAssetEntity, error)
}
