package repository

import (
	"asset_management_backend/common/service"
	"asset_management_backend/module/asset/domain/entity"
	"context"
)

type AssetRepository interface {
	// Tạo presigned url cho client upload file
	CreatePresignedUrls(ctx context.Context, fileNames []string) ([]service.PresignedResponse, error)
	// Random serial number ngẫu nhiên
	GenerateSerialNumber() string
	// Tạo ảnh label barcode ứng với serial number và upload, trả về image path
	GenerateLabelImage(ctx context.Context, serialNumber string) (*string, error)

	CreateAsset(ctx context.Context, request *entity.CreateAssetRequest) (*entity.AssetEntity, error)
	FindByID(ctx context.Context, assetID int) (*entity.AssetEntity, error)
	DeleteAssetFile(ctx context.Context, filePath string) error
}
