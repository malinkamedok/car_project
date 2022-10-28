package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type EngineerRepo struct {
	*postgres.Postgres
}

func NewEngineerRepo(pg *postgres.Postgres) *EngineerRepo {
	return &EngineerRepo{pg}
}

var _ usecase.EngineerRp = (*EngineerRepo)(nil)

func (er *EngineerRepo) GetEngineerByIdVendor(ctx context.Context, vendorID int64) ([]entity.Engineer, error) {
	query := `SELECT id, vendor_id, name, gender, experience, salary, factory_id FROM get_engineer_by_vendor($1)`

	rows, err := er.Pool.Query(ctx, query, vendorID)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var engineers []entity.Engineer

	for rows.Next() {
		var engineer entity.Engineer
		err = rows.Scan(
			&engineer.ID,
			&engineer.VendorId,
			&engineer.Name,
			&engineer.Gender,
			&engineer.Experience,
			&engineer.Salary,
			&engineer.FactoryID,
		)

		if err != nil {
			return nil, fmt.Errorf("error in parsing subsidy: %w", err)
		}

		engineers = append(engineers, engineer)
	}
	return engineers, nil
}
