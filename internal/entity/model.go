package entity

type Model struct {
	ID           int64  `json:"id"`
	VendorID     int64  `json:"vendor_id"`
	Name         string `json:"name"`
	WheelDrive   string `json:"wheeldrive"`
	Significance int64  `json:"significance"`
	Price        int64  `json:"price"`
	ProdCost     int64  `json:"prod_cost"`
	EngineerID   int64  `json:"engineer_id"`
	FactoryID    int64  `json:"factory_id"`
	Sales        int64  `json:"sales"`
}
