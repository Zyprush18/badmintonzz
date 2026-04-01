package domain

import "database/sql"

type Bookings struct {
	ID int
	Type_Payment string
	Status string
	User_id int
	Schedule_id int
	Created_at sql.NullTime
	Updated_at sql.NullTime
}