package domain

import "database/sql"


type Services struct {
	ID int
	Name string
	Price float64
	Created_at sql.NullTime
	Updated_at sql.NullTime
}