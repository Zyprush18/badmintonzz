package commands

import (
	"context"
	"crypto/rand"
	"strconv"
	"time"

	"github.com/Zyprush18/badmintonzz/internal/booking/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/booking/interfaces/request"
	"github.com/Zyprush18/badmintonzz/internal/config"
)

type CommandBooking interface {
	CreateBooking(ctx context.Context, booking *request.BookingRequest, user_id int, role string) (string, string, error)
	UpdateBooking(ctx context.Context, booking_id int, req *request.BookingUpdateRequest) error
	DeleteBooking(ctx context.Context, booking_id int) error
}

type repoBooking struct {
	repo infrastructure.RepoBooking
	midtrans config.MidtransCfg
}

func NewCommandsBooking(r infrastructure.RepoBooking, m config.MidtransCfg) CommandBooking {
	return &repoBooking{
		repo: r,
		midtrans: m,
	}
}


func (r *repoBooking) CreateBooking(ctx context.Context, booking *request.BookingRequest, user_id int, role string) (string, string,error) { 	
	data_svc, err := r.repo.GetPriceServices(ctx, booking.Service_id)
	if err != nil {
		return "", "", err
	}

	booking.Name_svc = data_svc.Name

	booking.Price = data_svc.Price

	booking.Order_Id = "badmintonzz-" + strconv.Itoa(booking.Service_id) + "-" + rand.Text()

	get_Midt, err := r.midtrans.SnapRequest(booking)
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
		Duration_Hour: booking.Hour,
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


func (r *repoBooking) UpdateBooking(ctx context.Context, booking_id int, req *request.BookingUpdateRequest) error {
	req.Booking_Id = booking_id
	req.Updated_At = time.Now()
	return r.repo.Update(ctx, req)
}


func (r *repoBooking) DeleteBooking(ctx context.Context, booking_id int) error {
	deleted_at := time.Now()
	return r.repo.Delete(ctx, booking_id, deleted_at)
}