package repository

import (
	"asset_management_backend/common/logger"
	"asset_management_backend/common/service"
	app_utils "asset_management_backend/common/utils"
	"asset_management_backend/infras"
	assetData "asset_management_backend/module/asset/data/data_source"
	"asset_management_backend/module/asset/data/model"
	"asset_management_backend/module/asset/domain/entity"
	labelTaskData "asset_management_backend/module/label_task/data/data_source"
	"context"
	"fmt"
	"image/png"
	"os"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/skip2/go-qrcode"
)

type assetRepositoryImpl struct {
	assetDS      assetData.AssetDataSource
	labelTaskDS  labelTaskData.LabelTaskDataSource
	minioService *service.FileStorageService
}

// Update implements AssetRepository.
func (a *assetRepositoryImpl) Update(ctx context.Context, assetID int, request entity.UpdateAssetRequest) error {

	if request.UpdateData != nil {
		if err := a.assetDS.UpdateByID(ctx, assetID, app_utils.StructToUpdateMap(request.UpdateData)); err != nil {
			return err
		}
	}

	if request.ShouldCreateNewLabel != nil && *request.ShouldCreateNewLabel {
		serialNumber, err := a.assetDS.GetSerialNumberByID(ctx, assetID)
		if err != nil {
			return err
		}
		newAssetPath, err := a.GenerateLabelImage(ctx, *serialNumber)
		if err == nil && newAssetPath != nil {
			updateMap := map[string]interface{}{
				"path":       *newAssetPath,
				"created_at": time.Now(),
			}
			if err := a.assetDS.UpdateLabelImageByAssetID(ctx, assetID, updateMap); err != nil {
				return err
			}
		}
	}

	if len(request.RemoveAssetFileIDs) != 0 {
		if err := a.assetDS.DeleteAssetFilesByIDs(ctx, request.RemoveAssetFileIDs); err != nil {
			return err
		}
	}

	if len(request.NewAssetFilePaths) != 0 {
		if _, err := a.assetDS.InsertAssetFiles(ctx, model.AssetFilesFromPath(request.NewAssetFilePaths, assetID)); err != nil {
			return err
		}
	}

	return nil
}

// FindByFilter implements AssetRepository.
func (a *assetRepositoryImpl) FindByFilter(ctx context.Context, filterQuery entity.AssetFilterQuery) ([]entity.AssetEntity, error) {
	filterModel := model.NewAssetFilterModelFromQuery(filterQuery)
	assetModels, err := a.assetDS.FindByFilter(ctx, filterModel)
	if err != nil {
		return nil, err
	}
	return model.AssetModelsToEntities(assetModels), nil
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
	assetLabelImage, err := a.assetDS.InsertLabelImage(ctx, model.AssetLabelImage{
		Path:      *request.AssetLabelPath,
		CreatedAt: time.Now(),
		AssetID:   assetModel.ID,
	})
	if err != nil {
		a.minioService.Delete(context.Background(), *request.AssetLabelPath)
		logger.Error("[AssetRepository]", "InsertAssetLabelImageErr", err)
	}
	if request.LabelTaskID != nil {
		updateData := map[string]interface{}{"asset_label_image_id": assetLabelImage.ID}
		a.labelTaskDS.UpdateByID(ctx, *request.LabelTaskID, updateData)
	}
	assetModel.AssetFiles = assetFiles
	assetModel.AssetLabelImage = assetLabelImage
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
		assetDS:      assetData.NewAssetDataSource(infras.GetDbProvider()),
		labelTaskDS:  labelTaskData.NewLabelTaskDataSource(infras.GetDbProvider()),
		minioService: service.NewFileStorageService(),
	}
}
