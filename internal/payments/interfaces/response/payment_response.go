package response

import (
	"database/sql"
	"time"
)

type PaymentResponse struct {
	ID             int       `db:"id" json:"id"`
	Order_Id	string		`db:"order_id" json:"order_id"`
	Amount         float32   `db:"amount" json:"amount"`
	Payment_Method sql.NullString `db:"payment_method" json:"payment_method"`
	Payment_Status string    `db:"payment_status" json:"payment_status"`
	Payment_Url    string    `db:"payment_url" json:"payment_url"`
	Transaction_id sql.NullString `db:"transaction_id" json:"transaction_id"`
	Created_at     time.Time `db:"created_at" json:"created_at,omitzero"`
	Updated_at     time.Time `db:"updated_at" json:"updated_at,omitzero"`
}
