package usecase

import (
	"context"
	"pahan/internal/entity"
)

type FactoryUseCase struct {
	repo FactoryRp
}

func NewFactoryUseCase(r FactoryRp) *FactoryUseCase {
	return &FactoryUseCase{
		repo: r,
	}
}

func (f *FactoryUseCase) GetFactoriesByVendor(ctx context.Context, vendorID int64) ([]entity.Factory, error) {
	return f.repo.GetFactoriesByVendorID(ctx, vendorID)
}
