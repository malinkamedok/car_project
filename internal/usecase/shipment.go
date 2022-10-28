package usecase

import (
	"context"
	"pahan/internal/entity"
)

type ShipmentUseCase struct {
	repo ShipmentRp
}

func NewShipmentUseCase(s ShipmentRp) *ShipmentUseCase {
	return &ShipmentUseCase{repo: s}
}

var _ Shipment = (*ShipmentUseCase)(nil)

func (s *ShipmentUseCase) CreateShipment(ctx context.Context, shipment entity.Shipment) error {
	return s.repo.CreateNewShipment(ctx, shipment)
}

func (s *ShipmentUseCase) GetShipments(ctx context.Context) ([]entity.Shipment, error) {
	return s.repo.GetAllShipments(ctx)
}
