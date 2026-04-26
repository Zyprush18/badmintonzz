package application

import (
	"github.com/Zyprush18/badmintonzz/internal/booking/application/commands"
	"github.com/Zyprush18/badmintonzz/internal/booking/application/queries"
	"github.com/Zyprush18/badmintonzz/internal/booking/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/config"
)

type BookingApplication interface {
	QueriesBooking() queries.QueriesBooking
	CommandsBooking() commands.CommandBooking
}

type repoBooking struct {
	r infrastructure.RepoBooking
	midtrans config.MidtransCfg
}


func NewApplicationBooking(r infrastructure.RepoBooking, m config.MidtransCfg) BookingApplication {
	return &repoBooking{
		r: r,
		midtrans: m,
	}
}

func (r *repoBooking) QueriesBooking() queries.QueriesBooking {
	return queries.NewQueriesBooking(r.r)
}

func (r *repoBooking) CommandsBooking() commands.CommandBooking {
	return commands.NewCommandsBooking(r.r, r.midtrans)
}