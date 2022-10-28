package entity

type Type struct {
	ID             int64  `json:"id"`
	Type           string `json:"type"`
	AdditionalInfo string `json:"additional_info"`
}
