package errs

import (
	"context"
	"database/sql"
)

var (
	ServerError string = "Internal Server Error"
	NotFoundRow error = sql.ErrNoRows
	InvalidRequest string = "Invalid request body"
	InvalidValidation string = "Invalid validation"
	RequestTimeout string = "Request Timeout"
	ContextTimeout error = context.DeadlineExceeded
)