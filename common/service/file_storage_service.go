package service

import (
	"asset_management_backend/common/logger"
	app_utils "asset_management_backend/common/utils"
	"asset_management_backend/infras"
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type FileStorageService struct {
	minioProvider *infras.MinioProvider
}

type PresignedResult struct {
	PresignedUrl string `json:"presigned_url"`
	FilePath     string `json:"file_path"`
	Url          string `json:"url"`
}

type PresignedResponse struct {
	FileName        string          `json:"file_name"`
	PresignedResult PresignedResult `json:"presigned_result"`
}

func (m *FileStorageService) CreatePresignedUrls(
	ctx context.Context,
	fileNames []string,
	folderName string,
	expires time.Duration,
) ([]PresignedResponse, error) {
	var presignedResponses []PresignedResponse
	for _, fileName := range fileNames {
		uniqueFileName := uuid.New().String()
		result, err := m.CreatePresignedUrl(
			ctx,
			fmt.Sprintf("%s/%s%s", folderName, uniqueFileName, filepath.Ext(fileName)),
			expires,
		)
		if err != nil {
			return nil, err
		}
		presignedResponses = append(presignedResponses, PresignedResponse{
			FileName:        fileName,
			PresignedResult: *result,
		})
	}
	return presignedResponses, nil
}

func (m *FileStorageService) CreatePresignedUrl(ctx context.Context, filePath string, expires time.Duration) (*PresignedResult, error) {
	presignedUrl, err := m.minioProvider.MinioClient.PresignedPutObject(
		ctx,
		m.minioProvider.Bucket,
		filePath,
		expires,
	)
	if err != nil {
		logger.Error("[MinioService] CreatePresignedUrlErr", err)
		return nil, err
	}
	result := &PresignedResult{
		PresignedUrl: presignedUrl.String(),
		FilePath:     filePath,
		Url:          fmt.Sprintf("%s://%s%s", presignedUrl.Scheme, presignedUrl.Host, presignedUrl.Path),
	}
	return result, nil
}

func (m *FileStorageService) Upload(ctx context.Context, objectName string, filePath string) (*string, error) {
	info, err := m.minioProvider.MinioClient.FPutObject(
		ctx,
		m.minioProvider.Bucket,
		objectName, filePath,
		minio.PutObjectOptions{ContentType: app_utils.GetContentType(filePath)},
	)
	if err != nil {
		return nil, err
	}
	return &info.Key, nil
}

func (m *FileStorageService) Delete(ctx context.Context, objectName string) error {
	err := m.minioProvider.MinioClient.RemoveObject(
		ctx,
		m.minioProvider.Bucket,
		objectName,
		minio.RemoveObjectOptions{ForceDelete: true},
	)
	if err != nil {
		return err
	}
	return nil
}

func NewFileStorageService() *FileStorageService {
	return &FileStorageService{
		minioProvider: infras.GetMinioProvider(),
	}
}
