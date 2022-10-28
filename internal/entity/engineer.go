package entity

type Engineer struct {
	ID         int64  `json:"id"`
	VendorId   int64  `json:"vendor_id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Experience int64  `json:"experience"`
	Salary     int64  `json:"salary"`
	FactoryID  int64  `json:"factory_id"`
}
