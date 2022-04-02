package models

type Location struct {
	Type        string    `json:"type,omitempty" validate:"required"`
	Coordinates []float64 `json:"coordinates,omitempty" validate:"required"`
}
