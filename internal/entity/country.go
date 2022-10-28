package entity

type Country struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	GdpUSD int64  `json:"gdp_usd"`
}
