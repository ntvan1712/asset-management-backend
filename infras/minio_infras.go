package infras

import (
	"asset_management_backend/app_config"
	"asset_management_backend/common/logger"
	app_utils "asset_management_backend/common/utils"
	"context"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioProvider *MinioProvider
var initMinioOnce sync.Once

type MinioProvider struct {
	MinioClient *minio.Client
	Bucket      string
}

func GetMinioProvider() *MinioProvider {
	initMinioOnce.Do(initMinioProvider)
	return minioProvider
}

func initMinioProvider() {
	config := app_config.GetAppConfig().MinioConfig

	// Initialize minio client object.
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: false, // http
	})
	if err != nil {
		logger.Fatal("[MinioInfras] init error", err)
	}
	minioProvider =  &MinioProvider{
		MinioClient: minioClient,
		Bucket:      config.Bucket,
	}

	logger.Info("[MinioInfras] Init MinioProvider")
}

func TestUpload() {
	objectName := "asset_labels/4.jpg"
	filePath := "4.jpg"

	info, err := GetMinioProvider().MinioClient.FPutObject(
		context.Background(), GetMinioProvider().Bucket, objectName, filePath,
		minio.PutObjectOptions{ContentType: app_utils.GetContentType(filePath)})
	if err != nil {
		logger.Fatal("[Upload]", err)
	}
	logger.Info("[UploadInfo]", info)
}

func TestDelete() {
	objectName := "4.jpg"
	// filePath := "4.jpg"

	err := GetMinioProvider().MinioClient.RemoveObject(
		context.Background(), GetMinioProvider().Bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		logger.Fatal("[Upload]", err)
	}
	// logger.Info("[UploadInfo]", info)
}
