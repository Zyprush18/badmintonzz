package domain

import "database/sql"


type Payments struct {
	ID int
	Amount float64
	Payment_Status string
	Payment_Method string
	Payment_Url sql.NullString
	Transaction_id sql.NullString
	Booking_id int
	Created_at sql.NullTime
	Updated_at sql.NullTime
}