package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type ShipmentRepo struct {
	*postgres.Postgres
}

func NewShipmentRepo(pg *postgres.Postgres) *ShipmentRepo {
	return &ShipmentRepo{pg}
}

var _ usecase.ShipmentRp = (*ShipmentRepo)(nil)

func (s *ShipmentRepo) CreateNewShipment(ctx context.Context, shipment entity.Shipment) error {
	query := `SELECT create_shipment($1, $2, $3)`

	rows, err := s.Pool.Query(ctx, query, shipment.OrderID, shipment.CountryToID, shipment.Date)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}

func (s *ShipmentRepo) GetAllShipments(ctx context.Context) ([]entity.Shipment, error) {
	query := `SELECT * FROM shipment`

	rows, err := s.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var shipments []entity.Shipment

	for rows.Next() {
		var shp entity.Shipment
		err = rows.Scan(&shp.ID,
			&shp.OrderID,
			&shp.CountryToID,
			&shp.Date,
			&shp.Cost)
		if err != nil {
			return nil, fmt.Errorf("cannot parsing shipment: %w", err)
		}
		shipments = append(shipments, shp)
	}
	return shipments, nil
}
