package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type SubsidyRepo struct {
	*postgres.Postgres
}

func NewSubsidyRepo(pg *postgres.Postgres) *SubsidyRepo {
	return &SubsidyRepo{pg}
}

var _ usecase.SubsidyRp = (*SubsidyRepo)(nil)

func (sr *SubsidyRepo) GetSubsidies(ctx context.Context) ([]entity.SubsidyCountry, error) {
	query := `SELECT * FROM get_all_subsidies()`

	rows, err := sr.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var subsidies []entity.SubsidyCountry

	for rows.Next() {
		var subsidy entity.SubsidyCountry
		err = rows.Scan(
			&subsidy.ID,
			&subsidy.CountryID,
			&subsidy.RequirePrice,
			&subsidy.RequiredWd,
			&subsidy.CountryName,
		)

		if err != nil {
			return nil, fmt.Errorf("error in parsing subsidy: %w", err)
		}

		subsidies = append(subsidies, subsidy)
	}
	return subsidies, nil
}

func (sr *SubsidyRepo) CreateAndLinkSubsidy(ctx context.Context, countryIDBy int64, requirePriceBy float64, requiredWdBy string) error {
	query := `SELECT create_subsidy($1, $2, $3)`

	rows, err := sr.Pool.Query(ctx, query, countryIDBy, requirePriceBy, requiredWdBy)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}

func (sr *SubsidyRepo) AcceptSubsidy(ctx context.Context, subsidyID int64, model entity.Model, componentEngineID, componentDoorID, componentBumperID, componentTransmissionID int64) error {
	query := `SELECT accept_subsidies($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`

	rows, err := sr.Pool.Query(ctx, query, subsidyID, model.VendorID, model.Name, model.Significance, model.EngineerID, model.FactoryID, componentEngineID, componentDoorID, componentBumperID, componentTransmissionID)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}

func (sr *SubsidyRepo) GetSubsidyByCountry(ctx context.Context, vendorID int64) ([]entity.Subsidy, error) {
	query := `SELECT * FROM get_subsidies_by_country_id($1)`

	rows, err := sr.Pool.Query(ctx, query, vendorID)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var subsidies []entity.Subsidy

	for rows.Next() {
		var subsidy entity.Subsidy
		err = rows.Scan(
			&subsidy.ID,
			&subsidy.CountryID,
			&subsidy.RequirePrice,
			&subsidy.RequiredWd,
		)

		if err != nil {
			return nil, fmt.Errorf("error in parsing subsidy: %w", err)
		}

		subsidies = append(subsidies, subsidy)
	}
	return subsidies, nil
}
