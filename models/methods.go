package models

type Method struct {
	Id               string `json:"_id,omitempty"`
	User             string `json:"user,omitempty" validate:"required"`
	Name             string `json:"name,omitempty" validate:"required"`
	Titular          string `json:"titular,omitempty" validate:"required"`
	FechaVencimiento string `json:"duedate,omitempty" validate:"required"`
	Number           string `json:"number,omitempty" validate:"required"`
	Type             string `json:"type,omitempty" validate:"required"`
	Sucursal         string `json:"sucursal,omitempty" validate:"required"`
}
