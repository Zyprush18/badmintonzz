package queries

import (
	"context"
	"database/sql"

	"github.com/Zyprush18/badmintonzz/internal/booking/domain"
	"github.com/Zyprush18/badmintonzz/internal/booking/infrastructure"
	booking "github.com/Zyprush18/badmintonzz/internal/booking/interfaces/response"
	payment "github.com/Zyprush18/badmintonzz/internal/payments/interfaces/response"
	svc_domain "github.com/Zyprush18/badmintonzz/internal/services/domain"
	user "github.com/Zyprush18/badmintonzz/internal/users/interfaces/response"
)

type QueriesBooking interface {
	GetBookings(ctx context.Context, user_id int, role string) ([]booking.BookingsResponse, error)
	GetBooking(ctx context.Context, id, user_id int, role string) (*booking.BookingsResponse, error)
}

type repoBooking struct {
	repo infrastructure.RepoBooking
}

func NewQueriesBooking(r infrastructure.RepoBooking) QueriesBooking {
	return &repoBooking{repo: r}
}

func (r *repoBooking) GetBookings(ctx context.Context, user_id int, role string) ([]booking.BookingsResponse, error) {
	var data []domain.Bookings
	var err error

	if role != "admin" {
		data, err = r.repo.GetBookingsByUserID(ctx, user_id)
	} else {
		data, err = r.repo.GetBookings(ctx)
	}

	if err != nil {
		return nil, err
	}

	bookings := r.ParseBookingSLice(data)

	return bookings, nil
}

func (r *repoBooking) GetBooking(ctx context.Context, id, user_id int, role string) (*booking.BookingsResponse, error) {
	var data *domain.Bookings
	var err error

	if role != "admin" {
		data, err = r.repo.GetBookingByUserIdAndId(ctx, user_id, id)
	} else {
		data, err = r.repo.GetBooking(ctx, id)
	}

	if err != nil {
		return nil, err
	}
	booking := r.ParseBooking(data)

	return &booking, nil
}


func (r *repoBooking) ParseBookingSLice(data []domain.Bookings) []booking.BookingsResponse {
	var bookings []booking.BookingsResponse
	for _, v := range data {
		bookings = append(bookings, r.ParseBooking(&v))
	}

	return bookings
}

func (r *repoBooking) ParseBooking(data *domain.Bookings) booking.BookingsResponse {
	return booking.BookingsResponse{
		ID:           data.ID_Booking,
		Date: data.Date,
		Start_Time: data.Start_Time,
		End_Time: data.End_Time,
		Type_Payment: data.Type_Payment,
		Status_Booking:       data.Status,
		Description: data.Description.String,
		User_id:      data.User_id,
		Payment_id:  data.Payment_id,
		Service_id: data.Service_id,
		Payment: payment.PaymentResponse{
			ID: data.Payment_id,
			Amount: data.Amount,
			Payment_Method: sql.NullString{String: data.Payment_Method, Valid: true},
			Payment_Status: data.Payment_Status,
			Payment_Url: data.Payment_Url.String,
			Transaction_id: sql.NullString{String: data.Transaction_id.String, Valid: true},
		},
		Service: svc_domain.Services{
			ID: data.Service_id,
			Name: data.Name_Service,
			Price: data.Price_Service,
		},
		User: user.UserResponse{
			ID:       data.User_id,
			Username: data.Username,
			Email:    data.Email,
			No_Hp:    data.No_Hp,
		},
		Created_at: data.CreatedAt_Booking,
		Updated_at: data.UpdatedAt_Booking,
	}
}
