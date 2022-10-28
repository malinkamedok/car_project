package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type ModelRepo struct {
	*postgres.Postgres
}

var _ usecase.ModelRp = (*ModelRepo)(nil)

func NewModelRepo(pg *postgres.Postgres) *ModelRepo {
	return &ModelRepo{pg}
}

func (m *ModelRepo) GetModels(ctx context.Context) ([]entity.ModelBig, error) {
	query := `SELECT * FROM get_all_models();`

	rows, err := m.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var models []entity.ModelBig

	for rows.Next() {
		var model entity.ModelBig
		err = rows.Scan(
			&model.ID,
			&model.VendorID,
			&model.Name,
			&model.WheelDrive,
			&model.Significance,
			&model.Price,
			&model.ProdCost,
			&model.EngineerID,
			&model.FactoryID,
			&model.Sales,
			&model.VendorName,
			&model.EngineerName,
			&model.CountryName,
		)

		if err != nil {
			return nil, fmt.Errorf("error in parsing model: %w", err)
		}

		models = append(models, model)
	}
	return models, nil
}

func (m *ModelRepo) DoNewModel(ctx context.Context, car entity.Model) error {
	query := `SELECT create_model($1, $2, $3, $4, $5, $6, $7)`

	rows, err := m.Pool.Query(ctx, query, car.VendorID, car.Name, car.WheelDrive, car.Significance, car.Price, car.EngineerID, car.FactoryID)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}
