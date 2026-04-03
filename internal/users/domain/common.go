package domain

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)


var (
	NotFoundUser string = "User not found"
	UserAlreadyExists string = "User already exists"
	ServerError string = "Internal Server Error"
	NotFoundRow error = sql.ErrNoRows
	InvalidID string = "Invalid User ID"
	InvalidRequest string = "Invalid request body"
	InvalidValidation string = "Invalid validation"
	DuplicateUser string = "User  email or phone number already exists"
)


func CheckDuplicate(err error) bool  {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == 1062 {
			return true
		}
	}
	return false
}