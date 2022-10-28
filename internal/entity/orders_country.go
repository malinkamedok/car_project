package entity

import "github.com/jackc/pgtype"

type OrdersCountry struct {
	VendorName   string             `json:"vendor_name"`
	ModelName    string             `json:"model_name"`
	ModelID      int64              `json:"model_id"`
	OrderID      int64              `json:"order_id"`
	Quantity     int64              `json:"quantity"`
	OrderType    string             `json:"order_type"`
	ShipmentCost int64              `json:"shipment_cost"`
	Date         pgtype.Timestamptz `json:"date"`
}
