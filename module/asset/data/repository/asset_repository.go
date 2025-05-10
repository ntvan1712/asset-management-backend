package repository

import (
	"asset_management_backend/common/service"
	sharedmodel "asset_management_backend/common/shared_model"
	"asset_management_backend/module/asset/domain/entity"
	"context"
)

type AssetRepository interface {
	// Tạo presigned url cho client upload file
	CreateAssetFilesPresignedUrls(ctx context.Context, fileNames []string) ([]service.PresignedResponse, error)
	// Random serial number ngẫu nhiên
	GenerateSerialNumber() string
	// Tạo ảnh label barcode ứng với serial number và upload, trả về image path
	GenerateLabelImage(ctx context.Context, serialNumber string) (*string, error)

	CreateAsset(ctx context.Context, request *entity.CreateAssetRequest) (*entity.AssetEntity, error)
	FindByID(ctx context.Context, assetID int) (*entity.AssetEntity, error)
	DeleteAssetFile(ctx context.Context, filePath string) error

	FindByFilter(
		ctx context.Context,
		filterQuery entity.AssetFilterQuery,
	) ([]entity.AssetEntity, error)

	Update(ctx context.Context, assetID int, request entity.UpdateAssetRequest) error
	FindAssetByBorrowerID(ctx context.Context, borrowerID int, paginateQuery sharedmodel.PaginateQuery) ([]entity.AssetEntity, error)
}
