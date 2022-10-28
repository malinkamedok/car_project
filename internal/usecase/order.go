package usecase

import (
	"context"
	"pahan/internal/entity"
)

type OrdersUseCase struct {
	repo OrderRp
}

var _ Order = (*OrdersUseCase)(nil)

func NewOrdersUseCase(r OrderRp) *OrdersUseCase {
	return &OrdersUseCase{repo: r}
}

func (o *OrdersUseCase) CreateOrder(ctx context.Context, order entity.Order, countryToID int64) error {
	return o.repo.CreateNewOrder(ctx, order, countryToID)
}

func (o *OrdersUseCase) GetOrders(ctx context.Context) ([]entity.Order, error) {
	return o.repo.GetAllOrders(ctx)
}

func (o *OrdersUseCase) GetOrdersByVendor(ctx context.Context, vendorID int64) ([]entity.OrdersVendor, error) {
	return o.repo.GetAllOrdersByVendor(ctx, vendorID)
}

func (o *OrdersUseCase) GetOrdersByCountry(ctx context.Context, countryID int64) ([]entity.OrdersCountry, error) {
	return o.repo.GetAllOrdersByCountry(ctx, countryID)
}

func (o *OrdersUseCase) DoOrder(ctx context.Context, orderID int64) error {
	return o.repo.DoOrder(ctx, orderID)
}
