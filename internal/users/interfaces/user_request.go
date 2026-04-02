package interfaces


type UserRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	No_Hp string `json:"phone"`
}