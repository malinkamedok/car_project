package repo

import (
	"context"
	"fmt"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type CountryRepo struct {
	*postgres.Postgres
}

func NewCountryRepo(pg *postgres.Postgres) *CountryRepo {
	return &CountryRepo{pg}
}

var _ usecase.CountryRp = (*CountryRepo)(nil)

func (c *CountryRepo) LoginCountry(ctx context.Context, s string) (int64, error) {
	query := `SELECT country.id FROM country where country.name = $1`

	rows, err := c.Pool.Query(ctx, query, s)
	if err != nil {
		return -1, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var countryID int64

	if rows.Next() {
		err = rows.Scan(&countryID)
		if err != nil {
			return -1, fmt.Errorf("error scan id country in login: %w", err)
		}
	} else {
		return -1, fmt.Errorf("сountry with this name does not exist")
	}

	return countryID, nil
}
