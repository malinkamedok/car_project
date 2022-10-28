package usecase

import (
	"context"
	"fmt"
	"pahan/internal/entity"
)

type VendorUseCase struct {
	repo VendorRp
}

func NewVendorUseCase(r VendorRp) *VendorUseCase {
	return &VendorUseCase{repo: r}
}

var _ Vendor = (*VendorUseCase)(nil)

func (v *VendorUseCase) GetAllVendors(ctx context.Context) ([]entity.Vendor, error) {
	listVendors, err := v.repo.GetVendors(ctx)
	if err != nil {
		return nil, fmt.Errorf("VendorUseCase - vendor list - s.repo.GetVendors: %w", err)
	}
	return listVendors, nil
}
