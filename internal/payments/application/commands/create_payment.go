package commands

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/config"
	"github.com/Zyprush18/badmintonzz/internal/payments/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/payments/interfaces/request"
)

type CommandPayment interface {
	UpdatePayment(ctx context.Context, order_id, transaction_status string) error
}

type repoPayment struct {
	repo infrastructure.PaymentRepo
	mdtrans config.MidtransCfg
}

func NewCommandsPayment(r infrastructure.PaymentRepo, mdtrans config.MidtransCfg) CommandPayment {
	return &repoPayment{
		repo: r,
		mdtrans: mdtrans,
	}
}

func (r *repoPayment) UpdatePayment(ctx context.Context, order_id, transaction_status string) error {
	data, err := r.mdtrans.GetData(order_id)
	if err != nil {
		return err
	}

	req := &request.PaymentRequest{
		Payment_method: data.PaymentType,
		Payment_status: r.CheckStatus(data.TransactionStatus),
		Transaction_id: data.TransactionID,
	}
	return r.repo.UpdatePayment(ctx, req, order_id)
}

func (r *repoPayment) CheckStatus(transaction_status string) string  {
	switch transaction_status {
	case "settlement":
		return "completed"
	case "pending":
		return "pending"
	case "expire":
		return "expired"
	case "refund":
		return "refunded"
	}
	return "failed"
}