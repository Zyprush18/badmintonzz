package request

import "time"

type BookingRequest struct {
	Name_svc string
	Hour int `json:"hour" validate:"required,number"`
	Amount float32 `db:"amount"`
	Payment_Method string `db:"payment_method" json:"payment_method" validate:"required"`
	Type_Payment string `db:"type_payment" json:"type_payment" validate:"required"`
	Status string `db:"status" json:"status" validate:"required"`
	User_id int `db:"user_id" json:"user_id"`
	Service_id int `db:"service_id" json:"service_id" validate:"required,number"`
	Updated_at time.Time `json:"updated_at"`
}


type BookingPaymentRequest struct {
	Amount float32 `db:"amount"`
	Type_Payment string `db:"type_payment"`
	
}
