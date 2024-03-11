package models

type RechargeResponse struct {
	Status   int      `json:"status,omitempty" validate:"required"`
	Recharge Recharge `json:"recharge,omitempty" validate:"required"`
}
