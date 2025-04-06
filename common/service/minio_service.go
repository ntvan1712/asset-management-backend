package service

import (
	"asset_management_backend/common/logger"
	app_utils "asset_management_backend/common/utils"
	"asset_management_backend/infras"
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
)

type MinioService struct {
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

func (m *MinioService) CreatePresignedUrl(ctx context.Context, filePath string, expires time.Duration) (*PresignedResult, error) {
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

func (m *MinioService) Upload(ctx context.Context, objectName string, filePath string) (*string, error) {
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

func (m *MinioService) Delete(ctx context.Context, objectName string) error {
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

func NewMinioService() *MinioService {
	return &MinioService{
		minioProvider: infras.GetMinioProvider(),
	}
}
