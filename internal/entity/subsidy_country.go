package entity

type SubsidyCountry struct {
	ID           int64  `json:"id"`
	CountryID    int64  `json:"country_id"`
	RequirePrice int64  `json:"require_price"`
	RequiredWd   string `json:"required_wd"`
	CountryName  string `json:"country_name"`
}
