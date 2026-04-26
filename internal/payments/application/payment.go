package application

import (
	"github.com/Zyprush18/badmintonzz/internal/config"
	_"github.com/Zyprush18/badmintonzz/internal/payments/application/commands"
	"github.com/Zyprush18/badmintonzz/internal/payments/application/queries"
	"github.com/Zyprush18/badmintonzz/internal/payments/infrastructure"
)


type ApplicationPayment interface {
	QueriesPayment() queries.QueriesPayment
}


type repoPayment struct {
	repo infrastructure.PaymentRepo
	midtrans config.MidtransCfg
}

func NewApplicationPayment(r infrastructure.PaymentRepo, m config.MidtransCfg) ApplicationPayment  {
	return &repoPayment{
		repo: r,
		midtrans: m,
	}
}

func (r *repoPayment) QueriesPayment() queries.QueriesPayment  {
	return queries.NewQueriesPayment(r.repo)
}