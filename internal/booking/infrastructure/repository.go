package infrastructure

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/booking/domain"
	"github.com/jmoiron/sqlx"
)

type RepoBooking interface {
	GetBookings(ctx context.Context) ([]domain.Bookings, error)
}

type database struct {
	db *sqlx.DB
}

func NewInfrastructureBooking(d *sqlx.DB) RepoBooking {
	return &database{db: d}
}

func (d *database) GetBookings(ctx context.Context) ([]domain.Bookings, error) {
	var bookings []domain.Bookings
	query := `
		SELECT
			b.id as booking_id,
			b.amount,
			b.type_payment,
			b.status,
			b.user_id,
			b.schedule_id,
			b.created_at as created_at_booking,
			b.updated_at as updated_at_booking,
			s.id as schedule_id,
			s.date as date_schedule,
			s.time as time_schedule,
			s.duration as duration_schedule,
			s.service_id,
			svc.name as name_service,
			svc.price as price_service,
			u.id as user_id,
			u.username,
			u.email,
			u.no_hp as phone
		FROM bookings as b
		LEFT JOIN schedules as s ON b.schedule_id = s.id
		LEFT JOIN users as u ON b.user_id = u.id
		LEFT JOIN services as svc ON s.service_id = svc.id;
	`

	if err := d.db.SelectContext(ctx, &bookings, query); err != nil {
		return nil, err
	}

	return bookings, nil
}
