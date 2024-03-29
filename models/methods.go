package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Method struct {
	Id               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	User             string             `json:"user,omitempty" validate:"required"`
	Name             string             `json:"name,omitempty" validate:"required"`
	Titular          string             `json:"titular,omitempty" validate:"required"`
	Duedate string             `json:"duedate,omitempty" validate:"required"`
	Number           string             `json:"number,omitempty" validate:"required"`
	Type             string             `json:"type,omitempty" validate:"required"`
	Sucursal         string             `json:"sucursal,omitempty" validate:"required"`
}
