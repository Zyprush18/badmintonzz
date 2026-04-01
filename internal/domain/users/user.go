package users

import "database/sql"

type Users struct {
	ID int
	Username string
	Email string
	Password string
	Phone string
	Created_at sql.NullTime
	Updated_at sql.NullTime
}