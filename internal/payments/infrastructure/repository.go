package infrastructure

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/payments/domain"
	"github.com/Zyprush18/badmintonzz/internal/payments/interfaces/request"
	"github.com/jmoiron/sqlx"
)

type PaymentRepo interface {
	GetAll(ctx context.Context) ([]domain.Payments, error)
	GetPaymentById(ctx context.Context, id int) (*domain.Payments, error)
	UpdatePayment(ctx context.Context, req *request.PaymentRequest, order_id string) error
}

type database struct {
	db *sqlx.DB
}

func NewRepositoryPayment(d *sqlx.DB) PaymentRepo {
	return &database{db: d}
}


func (d *database) GetAll(ctx context.Context) ([]domain.Payments, error)  {
	var data []domain.Payments
	query := `SELECT id, order_id, amount, payment_method, payment_status, payment_url, transaction_id, created_at, updated_at FROM payments WHERE deleted_at IS NULL`

	if err:= d.db.SelectContext(ctx, &data, query);err != nil {
		return nil, err
	}

	return data, nil
}

func (d *database) GetPaymentById(ctx context.Context, id int) (*domain.Payments, error) {
	query := `SELECT id, order_id, amount, payment_method, payment_status, payment_url, transaction_id, created_at, updated_at FROM payments WHERE id = ? AND deleted_at IS NULL`
	var data domain.Payments
	if err := d.db.GetContext(ctx, &data, query, id); err != nil {
		return nil, err
	}
	return &data, nil
}

func (d *database) UpdatePayment(ctx context.Context, req *request.PaymentRequest, order_id string) error {
	query := `UPDATE payments SET payment_method = ?, payment_status = ?, transaction_id = ? WHERE order_id = ?`
	_, err := d.db.ExecContext(ctx, query, req.Payment_method, req.Payment_status, req.Transaction_id, order_id)
	return err
}