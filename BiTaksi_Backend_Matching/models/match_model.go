package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MatchingRequest struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Type        string             `json:"type,omitempty" validate:"required"`
	Coordinates []float64          `json:"coordinates,omitempty" validate:"required"`
}
