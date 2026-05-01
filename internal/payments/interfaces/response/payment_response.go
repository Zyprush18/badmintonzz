package response

import (
	"time"
)

type PaymentResponse struct {
	ID             int       `json:"id"`
	Order_Id	string		`json:"order_id"`
	Amount         float32   `json:"amount"`
	Payment_Method string `json:"payment_method"`
	Payment_Status string    `json:"payment_status"`
	Payment_Url    string    `json:"payment_url"`
	Transaction_id string `json:"transaction_id"`
	Created_at     time.Time `json:"created_at,omitzero"`
	Updated_at     time.Time `json:"updated_at,omitzero"`
}
