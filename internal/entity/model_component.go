package entity

type ModelComponent struct {
	ID          int64 `json:"id"`
	ModelID     int64 `json:"model_id"`
	ComponentID int64 `json:"component_id"`
}
