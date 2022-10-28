package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type FactoryRepo struct {
	*postgres.Postgres
}

func NewFactoryRepo(pg *postgres.Postgres) *FactoryRepo {
	return &FactoryRepo{pg}
}

var _ usecase.FactoryRp = (*FactoryRepo)(nil)

func (f *FactoryRepo) GetFactoriesByVendorID(ctx context.Context, vendorID int64) ([]entity.Factory, error) {
	query := `SELECT id, vendor_id, max_workers, productivity FROM get_factory_by_vendor($1)`

	rows, err := f.Pool.Query(ctx, query, vendorID)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var factories []entity.Factory

	for rows.Next() {
		var fct entity.Factory
		err = rows.Scan(&fct.ID,
			&fct.VendorID,
			&fct.MaxWorkers,
			&fct.Productivity)
		if err != nil {
			return nil, fmt.Errorf("error in parsing factory: %w", err)
		}

		factories = append(factories, fct)
	}
	return factories, nil
}
