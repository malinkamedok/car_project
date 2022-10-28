package entity

type CountryTax struct {
	ID        int64 `json:"id"`
	CountryID int64 `json:"country_id"`
	Tax       int64 `json:"tax"`
}
