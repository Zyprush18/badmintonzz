package request


type UserAuthRegisterRequest struct {
	Username string `db:"username" json:"username" validate:"required,min=2,max=100"`
	Email string `db:"email" json:"email" validate:"required,email"`
	Password string `db:"password" json:"password" validate:"required,min=6,max=12"`
	Phone string `db:"no_hp" json:"phone" validate:"required,min=10,max=15"`
}
type UserAuthLoginRequest struct {
	Email string `db:"email" json:"email" validate:"required,email"`
	Password string `db:"password" json:"password" validate:"required,min=6,max=12"`
}
type UserRequest struct {
	Username string `db:"username" json:"username" validate:"required,min=2,max=100"`
	Email string `db:"email" json:"email" validate:"required,email"`
	Password string `db:"password" json:"password" validate:"required,min=6,max=12"`
	Phone string `db:"no_hp" json:"phone" validate:"required,min=10,max=15"`
	Role string `db:"role" json:"role" validate:"required,oneof=user admin"`
}