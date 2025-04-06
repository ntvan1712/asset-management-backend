package usecase

import (
	"asset_management_backend/common/service"
	"asset_management_backend/module/asset/domain/entity"
	"context"
)

type AssetUsecase interface {
	// Tạo presigned url cho client upload file
	GetPresignedUrls(ctx context.Context, fileNames []string) ([]service.PresignedResponse, error)

	CreateAsset(ctx context.Context, request *entity.CreateAssetRequest) (*entity.AssetEntity, error)
	GetAssetByID(ctx context.Context, assetID int) (*entity.AssetEntity, error)
}
