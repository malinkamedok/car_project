package entity

import "time"

type Shipment struct {
	ID          int64     `json:"id"`
	OrderID     int64     `json:"order_id"`
	CountryToID int64     `json:"country_to_id"`
	Date        time.Time `json:"date"`
	Cost        int64     `json:"cost"`
}
