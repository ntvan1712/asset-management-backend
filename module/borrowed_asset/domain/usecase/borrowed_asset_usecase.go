package usecase

import (
	"asset_management_backend/module/borrowed_asset/domain/entity"
	"context"
)

type BorrowedAssetUsecase interface {
	ReturnAsset(ctx context.Context, borrowedAssetId int) (*entity.BorrowedAssetEntity, error)
}
