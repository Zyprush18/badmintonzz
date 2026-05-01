package response

import (
	"time"

	payment "github.com/Zyprush18/badmintonzz/internal/payments/interfaces/response"
	"github.com/Zyprush18/badmintonzz/internal/services/domain"
	user "github.com/Zyprush18/badmintonzz/internal/users/interfaces/response"
)


type BookingsResponse struct {
	ID int `json:"id"`
	Date string `json:"date"`
	Start_Time string `json:"start_time"`
	End_Time string `json:"end_time"`
	Status_Booking string `json:"status_booking"`
	Description string `json:"description"`
	User_id int `json:"user_id"`
	Service_id int `json:"service_id"`
	Payment_id int `json:"payment_id"`
	User user.UserResponse `json:"user"`
	Service domain.Services `json:"service"`
	Payment  payment.PaymentResponse `json:"payment"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}