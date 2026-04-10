package domain

import (
	"time"
)

type Users struct {
	ID int	`json:"id"`
	Username string	`db:"username"`
	Email string	`db:"email"`
	Password string	`db:"password"`
	No_Phone string	`db:"no_hp"`
	Role string	`db:"role"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}