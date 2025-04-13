package usecase

import (
	"asset_management_backend/common/logger"
	"asset_management_backend/common/service"
	"asset_management_backend/module/asset/data/repository"
	"asset_management_backend/module/asset/domain/entity"
	"context"
)

type assetUsecaseImpl struct {
	assetRepo repository.AssetRepository
}

// GetAssetByID implements AssetUsecase.
func (a *assetUsecaseImpl) GetAssetByID(ctx context.Context, assetID int) (*entity.AssetEntity, error) {
	return a.assetRepo.FindByID(ctx, assetID)
}

// CreateAsset implements AssetUsecase.
func (a *assetUsecaseImpl) CreateAsset(ctx context.Context, request *entity.CreateAssetRequest) (*entity.AssetEntity, error) {
	if request.AssetLabelPath == nil {
		if request.SerialNumber == nil || *request.SerialNumber == "" {
			serial := a.assetRepo.GenerateSerialNumber()
			request.SerialNumber = &serial
		}
		uploadedAssetLabelPath, err := a.assetRepo.GenerateLabelImage(ctx, *request.SerialNumber)
		if err != nil {
			return nil, err
		}
		request.AssetLabelPath = uploadedAssetLabelPath
		logger.Info("[AssetUsecase]", "request.AssetLabelPath", request.AssetLabelPath)
	}
	asset, err := a.assetRepo.CreateAsset(ctx, request)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

// GetAssetFilesPresignedUrls implements AssetRepository.
func (a *assetUsecaseImpl) GetAssetFilesPresignedUrls(ctx context.Context, fileNames []string) ([]service.PresignedResponse, error) {
	return a.assetRepo.CreateAssetFilesPresignedUrls(ctx, fileNames)
}

func NewAssetUsecase() AssetUsecase {
	return &assetUsecaseImpl{
		assetRepo: repository.NewAssetRepository(),
	}
}
