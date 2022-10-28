package entity

type Storage struct {
	ID       int64 `json:"id"`
	ModelID  int64 `json:"model_id"`
	Amount   int64 `json:"amount"`
	VendorID int64 `json:"vendor_id"`
}
