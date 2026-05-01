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
	Status string `db:"status_booking"`
	Description sql.NullString `db:"description"`
	User_id int `db:"user_id"`
	Service_id int `db:"service_id"`
	Payment_id int `db:"payments_id"`
	Duration int `db:"duration_hour"`
	CreatedAt_Booking time.Time `db:"created_at_booking"`
	UpdatedAt_Booking time.Time `db:"updated_at_booking"`

	Order_id sql.NullString `db:"order_id"`
	Amount float32 `db:"amount"`
	Payment_Method sql.NullString `db:"payment_method"`
	Payment_Status string `db:"payment_status"`
	Payment_Url sql.NullString`db:"payment_url"`
	Transaction_id sql.NullString `db:"transaction_id"`
	CreatedAt_Payment time.Time `db:"created_at_payment"`
	UpdatedAt_Payment time.Time `db:"updated_at_payment"`


	Name_Service string `db:"name_service"`
	Price_Service float64 `db:"price_service"`
	CreatedAt_Service time.Time `db:"created_at_service"`
	UpdatedAt_Service time.Time `db:"updated_at_service"`


	Username string `db:"username"`
	Email string `db:"email"`
	No_Hp string `db:"phone"`
	CreatedAt_User time.Time `db:"created_at_user"`
	UpdatedAt_User time.Time `db:"updated_at_user"`
}


type GetService struct {
	Name string `db:"name"`
	Price float32 `db:"price"`
}