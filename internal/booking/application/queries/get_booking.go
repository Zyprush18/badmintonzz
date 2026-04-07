package queries

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/booking/infrastructure"
	booking "github.com/Zyprush18/badmintonzz/internal/booking/interfaces/response"
	schedule "github.com/Zyprush18/badmintonzz/internal/schedule/interfaces/response"
	user "github.com/Zyprush18/badmintonzz/internal/users/interfaces/response"
	"github.com/Zyprush18/badmintonzz/internal/services/domain"
)


type QueriesBooking interface {
	GetBookings(ctx context.Context) ([]booking.BookingsResponse, error)
}

type repoBooking struct {
	repo infrastructure.RepoBooking
}


func NewQueriesBooking(r infrastructure.RepoBooking) QueriesBooking  {
	return &repoBooking{repo: r}
}


func (r *repoBooking) GetBookings(ctx context.Context) ([]booking.BookingsResponse, error) {
	data, err := r.repo.GetBookings(ctx)
	if err != nil {
		return nil, err
	}

	var bookings []booking.BookingsResponse
	for _, v := range data {
		bookings = append(bookings, booking.BookingsResponse{
			ID: v.ID_Booking,
			Amount: v.Amount,
			Type_Payment: v.Type_Payment,
			Status: v.Status,
			User_id: v.User_id,
			Schedule_id: v.Schedule_id,
			Schedule: schedule.Schedules{
				ID: v.Schedule_id,
				Date: v.Date_Schedule,
				Time: v.Time_Schedule,
				Duration: v.Duration_Schedule,
				Service_id: v.Service_Id,
				Services: domain.Services{
					ID: v.Service_Id,
					Name: v.Name_Service,
					Price: v.Price_Service,
				},
			},
			User: user.UserResponse{
				ID: v.User_id,
				Username: v.Username,
				Email: v.Email,
				No_Hp: v.No_Hp,
			},
			Created_at: v.CreatedAt_Booking,
			Updated_at: v.UpdatedAt_Booking,
		})
	}

	return bookings, nil
}