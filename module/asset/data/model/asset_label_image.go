package model

import (
	"asset_management_backend/app_config"
	"asset_management_backend/module/asset/domain/entity"
	"context"
	"time"

	"github.com/uptrace/bun"
)

const TableAssetLabelImage = "asset_label_images"

type AssetLabelImage struct {
	bun.BaseModel `bun:"table:asset_label_images"`

	ID        int       `bun:",pk,autoincrement" json:"id"`
	Path      string    `bun:",notnull,type:varchar(100)" json:"path"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`

	AssetID int `json:"asset_id"`
}

var _ bun.AfterDeleteHook = (*AssetLabelImage)(nil)

func (*AssetLabelImage) AfterDelete(ctx context.Context, query *bun.DeleteQuery) error { return nil }

func (f *AssetLabelImage) ToEntity() *entity.AssetLabelImageEntity {
	if f == nil {
		return nil
	}
	return &entity.AssetLabelImageEntity{
		ID:        f.ID,
		Url:       app_config.GetAppConfig().MinioConfig.GetFullAssetUrl(f.Path),
		CreatedAt: f.CreatedAt,
	}
}
