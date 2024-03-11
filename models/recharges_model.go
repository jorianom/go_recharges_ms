package models

type Recharge struct {
	User   string `json:"user,omitempty" validate:"required"`
	Name   string `json:"amount,omitempty" validate:"required"`
	Method string `json:"method,omitempty" validate:"required"`
}
