package queries

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/payments/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/payments/interfaces/response"
)

type QueriesPayment interface {
	GetAll(ctx context.Context) ([]response.PaymentResponse, error)
}

type repoPayment struct {
	repo infrastructure.PaymentRepo
}

func NewQueriesPayment(r infrastructure.PaymentRepo) QueriesPayment  {
	return &repoPayment{
		repo: r,
	}
}


func (r *repoPayment) GetAll(ctx context.Context) ([]response.PaymentResponse, error)  {
	return r.repo.GetAll(ctx)
}