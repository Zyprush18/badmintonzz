package errs

import "database/sql"

var (
	ServerError string = "Internal Server Error"
	NotFoundRow error = sql.ErrNoRows
	InvalidRequest string = "Invalid request body"
	InvalidValidation string = "Invalid validation"
)