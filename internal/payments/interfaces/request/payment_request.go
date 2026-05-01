package request


type PaymentRequest struct {
	Payment_method string `db:"payment_method"`
	Payment_status string `db:"payment_status"`
	Transaction_id string `db:"transaction_id"`
}