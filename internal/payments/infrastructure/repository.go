package infrastructure

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/payments/interfaces/response"
	"github.com/jmoiron/sqlx"
)

type PaymentRepo interface {
	GetAll(ctx context.Context) ([]response.PaymentResponse, error)
}

type database struct {
	db *sqlx.DB
}

func NewRepositoryPayment(d *sqlx.DB) PaymentRepo {
	return &database{db: d}
}


func (d *database) GetAll(ctx context.Context) ([]response.PaymentResponse, error)  {
	var data []response.PaymentResponse
	query := `SELECT id, order_id, amount, payment_method, payment_status, payment_url, transaction_id, created_at, updated_at FROM payments WHERE deleted_at IS NULL`

	if err:= d.db.SelectContext(ctx, &data, query);err != nil {
		return nil, err
	}

	return data, nil
}
