package model

import (
	"asset_management_backend/app_config"
	app_utils "asset_management_backend/common/utils"
	"asset_management_backend/module/asset/domain/entity"
	"context"
	"time"

	"github.com/uptrace/bun"
)

type AssetFile struct {
	bun.BaseModel `bun:"table:asset_files"`

	ID        int       `bun:",pk,autoincrement" json:"id"`
	Path      string    `bun:",notnull,type:varchar(100)" json:"path"`
	FileType  string    `bun:",notnull,type:varchar(20)" json:"file_type"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`

	AssetID int `json:"asset_id"`
}

var _ bun.AfterDeleteHook = (*AssetFile)(nil)

func (*AssetFile) AfterDelete(ctx context.Context, query *bun.DeleteQuery) error { return nil }

func AssetFilesFromPath(paths []string, assetID int) []AssetFile {
	var assetFiles []AssetFile
	for _, path := range paths {
		assetFiles = append(assetFiles, AssetFile{
			Path:      path,
			FileType:  app_utils.GetFileTypeByPath(path),
			CreatedAt: time.Now(),
			AssetID:   assetID,
		})
	}
	return assetFiles
}

func (f *AssetFile) ToEntity() entity.AssetFileEntity {
	return entity.AssetFileEntity{
		ID:        f.ID,
		FileType:  f.FileType,
		Url:       app_config.GetAppConfig().MinioConfig.GetFullAssetUrl(f.Path),
		CreatedAt: f.CreatedAt,
	}
}

func AssetFileModelsToEntities(assetFiles []AssetFile) []entity.AssetFileEntity {
	var entities []entity.AssetFileEntity

	for _, location := range assetFiles {
		entities = append(entities, location.ToEntity())
	}

	return entities
}
