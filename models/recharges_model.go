package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recharge struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	User   string             `json:"user,omitempty" validate:"required"`
	Amount string             `json:"amount,omitempty" validate:"required"`
	Method string             `json:"method,omitempty" validate:"required"`
	Date   string             `json:"date,omitempty" validate:"required"`
	Status string             `json:"status,omitempty" validate:"required"`
	///fecha - statusS
}
