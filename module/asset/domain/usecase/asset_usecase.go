package usecase

import (
	"asset_management_backend/common/service"
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/module/asset/domain/entity"
	"context"
)

type AssetUsecase interface {
	// Tạo presigned url cho client upload file
	GetAssetFilesPresignedUrls(ctx context.Context, fileNames []string) ([]service.PresignedResponse, error)

	CreateAsset(ctx context.Context, request *entity.CreateAssetRequest) (*entity.AssetEntity, error)
	GetAssetByID(ctx context.Context, assetID int) (*entity.AssetEntity, error)

	SearchByFilter(
		ctx context.Context,
		filterQuery entity.AssetFilterQuery,
	) ([]entity.AssetEntity, error)

	Update(ctx context.Context, assetID int, request entity.UpdateAssetRequest) (*entity.AssetEntity, error)

	GetMyBorrowedAssets(ctx context.Context, borrowerID int, paginateQuery sharedmodel.PaginateQuery) ([]entity.AssetEntity, error)
}
