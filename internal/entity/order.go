package entity

type Order struct {
	ID        int64  `json:"id"`
	ModelID   int64  `json:"model_id"`
	Quantity  int64  `json:"quantity"`
	OrderType string `json:"order_type"`
}
