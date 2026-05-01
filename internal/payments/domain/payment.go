package domain

import (
	"database/sql"
	"time"
)


type Payments struct {
	ID int `db:"id"`
	Order_Id string `db:"order_id"`
	Amount float64 `db:"amount"`
	Payment_Status string `db:"payment_status"`
	Payment_Method sql.NullString `db:"payment_method"`
	Payment_Url sql.NullString `db:"payment_url"`
	Transaction_id sql.NullString `db:"transaction_id"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}