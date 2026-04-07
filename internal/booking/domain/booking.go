package domain

import "time"

type Bookings struct {
	ID_Booking int `db:"booking_id"`
	Amount float32 `db:"amount"`
	Type_Payment string `db:"type_payment" `
	Status string `db:"status"`
	User_id int `db:"user_id"`
	Schedule_id int `db:"schedule_id"`
	CreatedAt_Booking time.Time `db:"created_at_booking"`
	UpdatedAt_Booking time.Time `db:"updated_at_booking"`

	ID_Schedule int `db:"schedule_id"`
	Date_Schedule string `db:"date_schedule"`
	Time_Schedule string `db:"time_schedule"`
	Duration_Schedule int `db:"duration_schedule"`
	Service_Id int `db:"service_id"`
	Name_Service string `db:"name_service"`
	Price_Service float64 `db:"price_service"`

	ID_User int `db:"user_id"`
	Username string `db:"username"`
	Email string `db:"email"`
	No_Hp string `db:"phone"`

}