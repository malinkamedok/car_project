package usecase

import (
	"context"
	"pahan/internal/entity"
)

type ComponentUseCase struct {
	repo ComponentRp
}

func NewComponentUseCase(repo ComponentRp) *ComponentUseCase {
	return &ComponentUseCase{repo: repo}
}

var _ Component = (*ComponentUseCase)(nil)

func (c *ComponentUseCase) GetComponentsByVendorAndType(ctx context.Context, vendorID int64, typeID int64) ([]entity.Component, error) {
	return c.repo.GetComponentsByVendorIDAndTypeID(ctx, vendorID, typeID)
}

func (c *ComponentUseCase) GetComponents(ctx context.Context, typeComponent string) ([]entity.ComponentVendor, error) {
	return c.repo.GetAllComponents(ctx, typeComponent)
}

func (c *ComponentUseCase) CreateComponent(ctx context.Context, vendorID int64, typeID int64, name string, additionalInfo string) error {
	return c.repo.CreateCustomComponent(ctx, vendorID, typeID, name, additionalInfo)
}
