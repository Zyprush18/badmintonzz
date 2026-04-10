package response

import "time"


type UserResponse struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	No_Hp string `json:"phone"`
	Role string `json:"role,omitempty"`
	Created_at time.Time `json:"created_at,omitempty"`
	Updated_at time.Time `json:"updated_at,omitempty"`
}