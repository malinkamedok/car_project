package repo

import (
	"context"
	"fmt"
	"pahan/internal/entity"
	"pahan/internal/usecase"
	"pahan/pkg/postgres"
)

type OrderRepo struct {
	*postgres.Postgres
}

var _ usecase.OrderRp = (*OrderRepo)(nil)

func NewOrdersRepo(pg *postgres.Postgres) *OrderRepo {
	return &OrderRepo{pg}
}

func (o *OrderRepo) CreateNewOrder(ctx context.Context, ord entity.Order, countryToID int64) error {
	query := `SELECT create_order($1, $2, $3, $4)`

	rows, err := o.Pool.Query(ctx, query, ord.ModelID, ord.Quantity, ord.OrderType, countryToID)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}

func (o *OrderRepo) DoOrder(ctx context.Context, orderID int64) error {
	query := `SELECT do_order($1)`
	rows, err := o.Pool.Query(ctx, query, orderID)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	return nil
}

func (o *OrderRepo) GetAllOrders(ctx context.Context) ([]entity.Order, error) {
	query := `SELECT * FROM "order"`

	rows, err := o.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var orders []entity.Order

	for rows.Next() {
		var ord entity.Order
		err = rows.Scan(&ord.ID,
			&ord.ModelID,
			&ord.Quantity,
			&ord.OrderType)
		if err != nil {
			return nil, fmt.Errorf("error in parsing order: %w", err)
		}
		orders = append(orders, ord)
	}
	return orders, nil
}

func (o *OrderRepo) GetAllOrdersByVendor(ctx context.Context, vendorID int64) ([]entity.OrdersVendor, error) {
	query := `select * from get_orders_by_vendor_id($1)`

	rows, err := o.Pool.Query(ctx, query, vendorID)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var ordersVendor []entity.OrdersVendor

	for rows.Next() {
		var orderVendor entity.OrdersVendor
		err = rows.Scan(&orderVendor.ModelName,
			&orderVendor.ModelID,
			&orderVendor.CountryName,
			&orderVendor.OrderID,
			&orderVendor.Quantity,
			&orderVendor.OrderType,
			&orderVendor.ShipmentCost,
			&orderVendor.Date)
		if err != nil {
			return nil, fmt.Errorf("error in parsing order: %w", err)
		}
		ordersVendor = append(ordersVendor, orderVendor)
	}
	return ordersVendor, nil
}

func (o *OrderRepo) GetAllOrdersByCountry(ctx context.Context, countryID int64) ([]entity.OrdersCountry, error) {
	query := `select vendor_name , model_name , model_id, order_id , quantity, order_type, shipment_cost, shipment_date from get_orders_by_country_id($1)`

	rows, err := o.Pool.Query(ctx, query, countryID)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var ordersCountry []entity.OrdersCountry

	for rows.Next() {
		var orderCountry entity.OrdersCountry
		err = rows.Scan(
			&orderCountry.VendorName,
			&orderCountry.ModelName,
			&orderCountry.ModelID,
			&orderCountry.OrderID,
			&orderCountry.Quantity,
			&orderCountry.OrderType,
			&orderCountry.ShipmentCost,
			&orderCountry.Date)
		if err != nil {
			return nil, fmt.Errorf("error in parsing order: %w", err)
		}
		ordersCountry = append(ordersCountry, orderCountry)
	}
	return ordersCountry, nil
}
