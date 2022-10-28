package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type TypeRepo struct {
	*postgres.Postgres
}

func NewTypeRepo(pg *postgres.Postgres) *TypeRepo {
	return &TypeRepo{pg}
}

var _ usecase.TypeRp = (*TypeRepo)(nil)

func (t *TypeRepo) GetAllTypes(ctx context.Context) ([]entity.Type, error) {
	query := `SELECT id, type, additional_info FROM get_all_types()`

	rows, err := t.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var types []entity.Type

	for rows.Next() {
		var tp entity.Type
		err = rows.Scan(&tp.ID,
			&tp.Type,
			&tp.AdditionalInfo)
		if err != nil {
			return nil, fmt.Errorf("error in parsing type: %w", err)
		}
		types = append(types, tp)
	}
	return types, nil
}
