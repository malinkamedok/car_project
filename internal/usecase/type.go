package usecase

import (
	"context"
	"pahan/internal/entity"
)

type TypeUseCase struct {
	repo TypeRp
}

func NewTypeUseCase(repo TypeRp) *TypeUseCase {
	return &TypeUseCase{repo: repo}
}

var _ Type = (*TypeUseCase)(nil)

func (t *TypeUseCase) GetTypes(ctx context.Context) ([]entity.Type, error) {
	return t.repo.GetAllTypes(ctx)
}
