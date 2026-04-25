package request

import "time"

type BookingRequest struct {
	Name_svc string
	Order_Id string
	Start_Time string `json:"start_time" validate:"required"`
	Hour int `json:"hour" validate:"required,number"`
	Price float32 `db:"amount"`
	Service_id int `db:"service_id" json:"service_id" validate:"required,number"`
}


type BookingPaymentRequest struct {
	Day string `db:"day"`
	Start_Time string `db:"start_time"`
	End_Time string `db:"end_time"`
	
	Order_ID string `db:"order_id"`
	Amount float32 `db:"amount"`
	Payment_Url string `db:"payment_url"`
	Created_At_Payment time.Time `db:"created_at_payment"`
	
	Date string `db:"date"`
	Start_Time_Booking string `db:"start_time_booking"`
	End_Time_Booking string `db:"end_time_booking"`
	Status_Booking string `db:"status_booking"`
	Users_Id int `db:"user_id"`
	Service_Id int `db:"service_id"`
	Payment_ID int `db:"payment_id"`
	Created_At_Booking time.Time `db:"created_at_booking"`
}
