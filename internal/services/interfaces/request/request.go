package request

type Services struct {
	Name string `db:"name" json:"name" validate:"required,min=2"`
	Price float64 `db:"price" json:"price" validate:"required,numeric"`
}