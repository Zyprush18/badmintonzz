package domain

import (
	"database/sql"
	"time"
)

type Bookings struct {
	ID_Booking int `db:"booking_id"`
	Date string `db:"date"`
	Start_Time string `db:"start_time"`
	End_Time string `db:"end_time"`
	Type_Payment string `db:"type_payment" `
	Status string `db:"status_booking"`
	Description sql.NullString `db:"description"`
	User_id int `db:"user_id"`
	Service_id int `db:"service_id"`
	Payment_id int `db:"payments_id"`
	CreatedAt_Booking time.Time `db:"created_at_booking"`
	UpdatedAt_Booking time.Time `db:"updated_at_booking"`

	Amount float32 `db:"amount"`
	Payment_Method string `db:"payment_method"`
	Payment_Status string `db:"payment_status"`
	Payment_Url sql.NullString`db:"payment_url"`
	Transaction_id sql.NullString `db:"transaction_id"`


	Name_Service string `db:"name_service"`
	Price_Service float64 `db:"price_service"`


	Username string `db:"username"`
	Email string `db:"email"`
	No_Hp string `db:"phone"`

}


type GetService struct {
	Name string `db:"name"`
	Price float32 `db:"price"`
}