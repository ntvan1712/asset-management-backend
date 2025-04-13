package repository

import (
	"asset_management_backend/common/logger"
	"asset_management_backend/common/service"
	"asset_management_backend/infras"
	datasource "asset_management_backend/module/asset/data/data_source"
	"asset_management_backend/module/asset/data/model"
	"asset_management_backend/module/asset/domain/entity"
	"context"
	"fmt"
	"image/png"
	"os"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/skip2/go-qrcode"
)

type assetRepositoryImpl struct {
	assetDS      datasource.AssetDataSource
	minioService *service.FileStorageService
}

// FindByID implements AssetRepository.
func (a *assetRepositoryImpl) FindByID(ctx context.Context, assetID int) (*entity.AssetEntity, error) {
	assetModel, err := a.assetDS.FindByID(ctx, assetID)
	if err != nil {
		return nil, err
	}
	return assetModel.ToEntity(), nil

}

// DeleteAssetFile implements AssetRepository.
func (a *assetRepositoryImpl) DeleteAssetFile(ctx context.Context, filePath string) error {
	panic("unimplemented")
}

// CreateAsset implements AssetRepository.
func (a *assetRepositoryImpl) CreateAsset(ctx context.Context, request *entity.CreateAssetRequest) (*entity.AssetEntity, error) {
	newAssetModel := model.NewAssetModelFromRequest(request)
	assetModel, err := a.assetDS.Insert(ctx, newAssetModel)
	if err != nil {
		// Remove asset file
		go func() {
			a.minioService.Delete(context.Background(), *request.AssetLabelPath)
			for _, filePath := range request.AssetFilePaths {
				a.minioService.Delete(context.Background(), filePath)
			}
		}()
		return nil, err
	}
	assetFiles, err := a.assetDS.InsertAssetFiles(ctx, model.AssetFilesFromPath(request.AssetFilePaths, assetModel.ID))
	if err != nil {
		go func() {
			for _, filePath := range request.AssetFilePaths {
				a.minioService.Delete(context.Background(), filePath)
			}
		}()
		logger.Error("[AssetRepository]", "InsertAssetFilesErr", err)
	}
	assetModel.AssetFiles = assetFiles
	return assetModel.ToEntity(), nil
}

// GenerateLabelImage implements AssetRepository.
func (a *assetRepositoryImpl) GenerateLabelImage(ctx context.Context, serialNumber string) (*string, error) {
	qrCode, err := qrcode.New(serialNumber, qrcode.Medium)
	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s.png", serialNumber)
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := png.Encode(file, qrCode.Image(256)); err != nil {
		return nil, err
	}

	objectName := fmt.Sprintf("labels/%s", fileName)
	uploadedKey, err := a.minioService.Upload(ctx, objectName, fileName)
	if err != nil {
		return nil, err
	}
	go func() {
		if err := os.Remove(fileName); err != nil {
			logger.Error("[AssetRepository]", "GenerateLabelImageErr", err)
		}
	}()
	return uploadedKey, nil
}

// GenerateSerialNumber implements AssetRepository.
func (a *assetRepositoryImpl) GenerateSerialNumber() string {
	const alphabet = "ABCDEFGHIJKLMNPQRSTUVWXYZ123456789"
	id, _ := gonanoid.Generate(alphabet, 10)
	return id
}

// CreateAssetFilesPresignedUrls implements AssetRepository.
func (a *assetRepositoryImpl) CreateAssetFilesPresignedUrls(
	ctx context.Context,
	fileNames []string,
) ([]service.PresignedResponse, error) {
	return a.minioService.CreatePresignedUrls(ctx, fileNames, "asset-files", time.Minute*5)
}

func NewAssetRepository() AssetRepository {
	return &assetRepositoryImpl{
		assetDS:      datasource.NewAssetDataSource(infras.GetDbProvider()),
		minioService: service.NewFileStorageService(),
	}
}
