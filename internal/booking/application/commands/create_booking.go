package commands

import (
	"context"
	"crypto/rand"
	"strconv"
	"time"

	"github.com/Zyprush18/badmintonzz/internal/booking/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/booking/interfaces/request"
)

type CommandBooking interface {
	CreateBooking(ctx context.Context, booking *request.BookingRequest, user_id int, role string) (string, string, error)
}

type repoBooking struct {
	repo infrastructure.RepoBooking
}

func NewCommandsBooking(r infrastructure.RepoBooking) CommandBooking {
	return &repoBooking{repo: r}
}


func (r *repoBooking) CreateBooking(ctx context.Context, booking *request.BookingRequest, user_id int, role string) (string, string,error) { 	
	data_svc, err := r.repo.GetPriceServices(ctx, booking.Service_id)
	if err != nil {
		return "", "", err
	}

	booking.Name_svc = data_svc.Name

	booking.Price = data_svc.Price

	booking.Order_Id = "badmintonzz-" + strconv.Itoa(booking.Service_id) + "-" + rand.Text()

	mid := NewMidtrans(booking)

	get_Midt, err := mid.SnapRequest()
	if err != nil {
		return "", "", err
	}

	time_app := time.Now()

	start_time, err := time.Parse("15:04:05", booking.Start_Time)
	if err != nil {
		return "", "", err
	}
	end_time := start_time.Add(time.Duration(booking.Hour) * time.Hour)
	
	day := time_app.Weekday().String()
	bookingPayment := &request.BookingPaymentRequest{
		Day: day,
		Start_Time: start_time.Format(time.TimeOnly),
		End_Time: end_time.Format(time.TimeOnly),

		Order_ID: booking.Order_Id,
		Amount: booking.Price * float32(booking.Hour),
		Payment_Url: get_Midt.RedirectURL,
		Created_At_Payment: time_app,
	
		Date: time_app.Format(time.DateOnly),
		Start_Time_Booking: start_time.Format(time.TimeOnly),
		End_Time_Booking: end_time.Format(time.TimeOnly),
		Status_Booking: "pending",
		Users_Id: user_id,
		Service_Id: booking.Service_id,
		Created_At_Booking: time_app,
	}

	if err:= r.repo.CreateBooking(ctx,bookingPayment);err != nil {
		return "", "", err
	}


	return get_Midt.Token, get_Midt.RedirectURL, nil
}