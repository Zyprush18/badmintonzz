package application

import (
	"github.com/Zyprush18/badmintonzz/internal/booking/application/commands"
	"github.com/Zyprush18/badmintonzz/internal/booking/application/queries"
	"github.com/Zyprush18/badmintonzz/internal/booking/infrastructure"
)

type BookingApplication interface {
	QueriesBooking() queries.QueriesBooking
	CommandsBooking() commands.CommandBooking
}

type repoBooking struct {
	r infrastructure.RepoBooking
}


func NewApplicationBooking(r infrastructure.RepoBooking) BookingApplication {
	return &repoBooking{r: r}
}

func (r *repoBooking) QueriesBooking() queries.QueriesBooking {
	return queries.NewQueriesBooking(r.r)
}

func (r *repoBooking) CommandsBooking() commands.CommandBooking {
	return commands.NewCommandsBooking(r.r)
}