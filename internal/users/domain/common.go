package domain

import "database/sql"


var (
	NotFoundUser string = "User not found"
	UserAlreadyExists string = "User already exists"
	ServerError string = "Internal Server Error"
	NotFoundRow error = sql.ErrNoRows
	InvalidID string = "Invalid User ID"
)