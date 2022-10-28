package usecase

import (
	"context"
	"fmt"
	"pahan/internal/entity"
)

type EngineerUseCase struct {
	repo EngineerRp
}

func NewEngineerUseCase(r EngineerRp) *EngineerUseCase {
	return &EngineerUseCase{repo: r}
}

func (en *EngineerUseCase) GetAllEngineerByIdVendor(ctx context.Context, vendorID int64) ([]entity.Engineer, error) {
	listEngineers, err := en.repo.GetEngineerByIdVendor(ctx, vendorID)
	if err != nil {
		return nil, fmt.Errorf("EngineerUseCase - Engineer list - m.repo.GetAllEngineerByIdVendor: %w", err)
	}
	return listEngineers, nil
}
