package entity

type Component struct {
	ID             int64  `json:"id"`
	VendorID       int64  `json:"vendor_id"`
	TypeID         int64  `json:"type_id"`
	Name           string `json:"name"`
	AdditionalInfo string `json:"additional_info"`
}
