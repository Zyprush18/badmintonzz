package response

import (
	"time"

	schedule "github.com/Zyprush18/badmintonzz/internal/schedule/interfaces/response"
	user "github.com/Zyprush18/badmintonzz/internal/users/interfaces/response"
)


type BookingsResponse struct {
	ID int `json:"id"`
	Amount float32 `json:"amount"`
	Type_Payment string `json:"type_payment"`
	Status string `json:"status"`
	User_id int `json:"user_id"`
	Schedule_id int `json:"schedule_id"`
	Schedule schedule.Schedules `json:"schedule"`
	User user.UserResponse `json:"user"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}