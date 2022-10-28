package usecase

import (
	"context"
	"fmt"
	"pahan/internal/entity"
)

type ModelUseCase struct {
	repo ModelRp
}

func NewModelUseCase(r ModelRp) *ModelUseCase {
	return &ModelUseCase{repo: r}
}

var _ Model = (*ModelUseCase)(nil)

func (m *ModelUseCase) GetAllModels(ctx context.Context) ([]entity.ModelBig, error) {
	listModels, err := m.repo.GetModels(ctx)
	if err != nil {
		return nil, fmt.Errorf("ModelUseCase - model list - m.repo.GetModels: %w", err)
	}
	return listModels, nil
}

func (uc *ModelUseCase) NewModel(ctx context.Context, car entity.Model) error {
	err := uc.repo.DoNewModel(ctx, car)
	if err != nil {
		return fmt.Errorf("DesignUseCase - NewModel - s.repo.DoNewModel: %w", err)
	}
	return nil
}
