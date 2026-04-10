package commands

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/booking/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/booking/interfaces/request"
)

type CommandBooking interface {
	CreateBooking(ctx context.Context, booking *request.BookingRequest, user_id int, role string) error
}

type repoBooking struct {
	repo infrastructure.RepoBooking
}

func NewCommandsBooking(r infrastructure.RepoBooking) CommandBooking {
	return &repoBooking{repo: r}
}


func (r *repoBooking) CreateBooking(ctx context.Context, booking *request.BookingRequest, user_id int, role string) error {
	if role != "admin" {
		booking.User_id = user_id
	}
	return r.repo.CreateBooking(ctx, booking)
}