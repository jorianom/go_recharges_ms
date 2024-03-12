package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RechargeResponse struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Status   int                `json:"status,omitempty" validate:"required"`
	Recharge Recharge           `json:"recharge,omitempty" validate:"required"`
}

type MethodResponse struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Message string             `json:"message,omitempty"`
	Status  int                `json:"status,omitempty" validate:"required"`
	Method  Method             `json:"method,omitempty" validate:"required"`
}
