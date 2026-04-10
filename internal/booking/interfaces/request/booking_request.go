package request

import "time"

type BookingRequest struct {
	Amount float32 `db:"amount" json:"amount" validate:"required,number"`
	Type_Payment string `db:"type_payment" json:"type_payment" validate:"required"`
	Status string `db:"status" json:"status" validate:"required"`
	User_id int `db:"user_id" json:"user_id"`
	Schedule_id int `db:"schedule_id" json:"schedule_id" validate:"required,number"`
	Updated_at time.Time `json:"updated_at"`
}