package domain

import "time"


type Services struct {
	ID int
	Name string
	Price float64
	Created_at time.Time
	Updated_at time.Time
}