package entity

type Factory struct {
	ID           int64 `json:"id"`
	VendorID     int64 `json:"vendor_id"`
	MaxWorkers   int64 `json:"max_workers"`
	Productivity int64 `json:"productivity"`
}
