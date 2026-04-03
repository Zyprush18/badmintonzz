package domain

import (
	"time"
)

type Users struct {
	ID int	`json:"id"`
	Username string	`json:"username" db:"username"`
	Email string	`json:"email" db:"email"`
	Password string	`json:"password" db:"password"`
	No_Phone string	`json:"phone" db:"no_hp"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}