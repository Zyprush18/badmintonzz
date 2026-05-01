package queries

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/payments/domain"
	"github.com/Zyprush18/badmintonzz/internal/payments/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/payments/interfaces/response"
)

type QueriesPayment interface {
	GetAll(ctx context.Context) ([]response.PaymentResponse, error)
	GetPaymentById(ctx context.Context, id int) (*response.PaymentResponse, error)
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
	data, err := r.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []response.PaymentResponse
	for _, payment := range data {
		responses = append(responses, *r.ParsePaymentToResponse(&payment))
	}
	return responses, nil
}

func (r *repoPayment) GetPaymentById(ctx context.Context, id int) (*response.PaymentResponse, error) {
	payment, err := r.repo.GetPaymentById(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.ParsePaymentToResponse(payment), nil
}

func (r *repoPayment) ParsePaymentToResponse(payment *domain.Payments) *response.PaymentResponse {
	return &response.PaymentResponse{
		ID:             payment.ID,
		Order_Id:       payment.Order_Id,
		Amount:         float32(payment.Amount),
		Payment_Method: payment.Payment_Method.String,
		Payment_Status: payment.Payment_Status,
		Payment_Url:    payment.Payment_Url.String,
		Transaction_id: payment.Transaction_id.String,
		Created_at:     payment.Created_at,
		Updated_at:     payment.Updated_at,
	}
}