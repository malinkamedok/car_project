package entity

type ComponentVendor struct {
	ID             int64  `json:"id"`
	VendorID       int64  `json:"vendor_id"`
	VendorName     string `json:"vendor_name"`
	TypeID         int64  `json:"type_id"`
	Name           string `json:"name"`
	AdditionalInfo string `json:"additional_info"`
}
