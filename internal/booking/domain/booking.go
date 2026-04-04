package domain

import "database/sql"

type Bookings struct {
	ID int
	Amount float32
	Type_Payment string
	Status string
	User_id int
	Schedule_id int
	Created_at sql.NullTime
	Updated_at sql.NullTime
}