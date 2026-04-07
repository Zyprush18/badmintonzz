package application

import (
	"github.com/Zyprush18/badmintonzz/internal/booking/application/queries"
	"github.com/Zyprush18/badmintonzz/internal/booking/infrastructure"
)

type BookingApplication interface {
	QueriesBooking() queries.QueriesBooking
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