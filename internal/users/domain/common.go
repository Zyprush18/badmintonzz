package domain

import (

	"github.com/go-sql-driver/mysql"
)


var (
	NotFoundUser string = "User not found"
	UserAlreadyExists string = "User already exists"
	InvalidID string = "Invalid User ID"
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