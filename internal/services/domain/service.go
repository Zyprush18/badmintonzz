package domain

import "time"


type Services struct {
	ID int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Price float64 `db:"price" json:"price"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	Updated_at time.Time `db:"updated_at" json:"updated_at"`
}