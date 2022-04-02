package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DriverLocation struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Location Location           `json:"location,omitempty" validate:"required"`
}
