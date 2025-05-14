package usecase

import (
	"asset_management_backend/module/borrowed_asset/data/repository"
	"asset_management_backend/module/borrowed_asset/domain/entity"
	"context"
)

type borrowedAssetUsecaseImpl struct {
	borrowedAssetRepo repository.BorrowedAssetRepository
}

// ReturnAsset implements BorrowedAssetRepository.
func (b borrowedAssetUsecaseImpl) ReturnAsset(ctx context.Context, borrowedAssetId int) (*entity.BorrowedAssetEntity, error) {
	return b.borrowedAssetRepo.ReturnAsset(ctx, borrowedAssetId)

}

func NewBorrowedAssetUsecase() BorrowedAssetUsecase {
	return borrowedAssetUsecaseImpl{
		borrowedAssetRepo: repository.NewBorrowedAssetRepository(),
	}
}
