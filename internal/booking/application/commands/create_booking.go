package commands

import (
	"context"
	"log"

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

	data_svc, err := r.repo.GetPriceServices(ctx, booking.Service_id)
	if err != nil {
		return err
	}

	booking.Name_svc = data_svc.Name

	booking.Amount = data_svc.Price

	mid := NewMidtrans(booking)

	get_Midt, err := mid.SnapRequest()
	if err != nil {
		return err
	}



	log.Println(get_Midt)


	return nil
}