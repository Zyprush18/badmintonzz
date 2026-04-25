package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Zyprush18/badmintonzz/internal/booking/domain"
	"github.com/Zyprush18/badmintonzz/internal/booking/interfaces/request"
	"github.com/jmoiron/sqlx"
)

type RepoBooking interface {
	GetBookings(ctx context.Context) ([]domain.Bookings, error)
	GetBooking(ctx context.Context, id int) (*domain.Bookings, error)
	GetBookingsByUserID(ctx context.Context, userID int) ([]domain.Bookings, error)
	GetBookingByUserIdAndId(ctx context.Context, userId int, bookingId int) (*domain.Bookings, error)
	CreateBooking(ctx context.Context, booking *request.BookingPaymentRequest) error
	GetPriceServices(ctx context.Context, serviceID int) (*domain.GetService, error)
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
			b.date,
			b.start_time,
			b.end_time,
			b.type_payment,
			b.status_booking,
			b.created_at as created_at_booking,
			b.updated_at as updated_at_booking,

			p.id as payments_id,
			p.amount,
			p.payment_method,
			p.payment_status,
			p.payment_url,
			p.transaction_id,
			
			svc.id as service_id,
			svc.name as name_service,
			svc.price as price_service,

			u.id as user_id,
			u.username,
			u.email,
			u.no_hp as phone
		FROM bookings as b
		LEFT JOIN payments as p ON b.payments_id = p.id
		LEFT JOIN users as u ON b.user_id = u.id
		LEFT JOIN services as svc ON b.service_id = svc.id;
	`

	if err := d.db.SelectContext(ctx, &bookings, query); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (d *database) GetBooking(ctx context.Context, id int) (*domain.Bookings, error) {
	var booking domain.Bookings
	query := `
		SELECT
			b.id as booking_id,
			b.date,
			b.start_time,
			b.end_time,
			b.type_payment,
			b.status_booking,
			b.created_at as created_at_booking,
			b.updated_at as updated_at_booking,

			p.id as payments_id,
			p.amount,
			p.payment_method,
			p.payment_status,
			p.payment_url,
			p.transaction_id,
			
			svc.id as service_id,
			svc.name as name_service,
			svc.price as price_service,

			u.id as user_id,
			u.username,
			u.email,
			u.no_hp as phone
		FROM bookings as b
		LEFT JOIN payments as p ON b.payments_id = p.id
		LEFT JOIN users as u ON b.user_id = u.id
		LEFT JOIN services as svc ON b.service_id = svc.id
		WHERE b.id = ?
	`
	if err := d.db.GetContext(ctx, &booking, query, id); err != nil {
		return nil, err
	}

	return &booking, nil
}

func (d *database) GetBookingsByUserID(ctx context.Context, userID int) ([]domain.Bookings, error) {
	var bookings []domain.Bookings
	query := `
		SELECT
			b.id as booking_id,
			b.date,
			b.start_time,
			b.end_time,
			b.type_payment,
			b.status_booking,
			b.created_at as created_at_booking,
			b.updated_at as updated_at_booking,

			p.id as payments_id,
			p.amount,
			p.payment_method,
			p.payment_status,
			p.payment_url,
			p.transaction_id,
			
			svc.id as service_id,
			svc.name as name_service,
			svc.price as price_service,

			u.id as user_id,
			u.username,
			u.email,
			u.no_hp as phone
		FROM bookings as b
		LEFT JOIN payments as p ON b.payments_id = p.id
		LEFT JOIN users as u ON b.user_id = u.id
		LEFT JOIN services as svc ON b.service_id = svc.id
		WHERE b.user_id = ?
	`

	if err := d.db.SelectContext(ctx, &bookings, query, userID); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (d *database) GetBookingByUserIdAndId(ctx context.Context, userId int, bookingId int) (*domain.Bookings, error) {
	var booking domain.Bookings
	query := `
		SELECT
			b.id as booking_id,
			b.date,
			b.start_time,
			b.end_time,
			b.type_payment,
			b.status_booking,
			b.created_at as created_at_booking,
			b.updated_at as updated_at_booking,

			p.id as payments_id,
			p.amount,
			p.payment_method,
			p.payment_status,
			p.payment_url,
			p.transaction_id,
			
			svc.id as service_id,
			svc.name as name_service,
			svc.price as price_service,

			u.id as user_id,
			u.username,
			u.email,
			u.no_hp as phone
		FROM bookings as b
		LEFT JOIN payments as p ON b.payments_id = p.id
		LEFT JOIN users as u ON b.user_id = u.id
		LEFT JOIN services as svc ON b.service_id = svc.id
		WHERE b.user_id = ? AND b.id = ?
	`
	if err := d.db.GetContext(ctx, &booking, query, userId, bookingId); err != nil {
		return nil, err
	}

	return &booking, nil
}

func (d *database) CreateBooking(ctx context.Context, booking *request.BookingPaymentRequest) error {
	tx, err := d.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	defer func ()  {
		if err != nil {
			tx.Rollback()
		}
	}()

	var result_count_BussHour int

	query_cek_bussHour := `SELECT count(*) FROM bussiness_hour WHERE day=? AND start_time <= ? AND end_time >= ? AND is_open = TRUE;`

	if err:= tx.QueryRowContext(ctx, query_cek_bussHour, booking.Day, booking.Start_Time, booking.End_Time).Scan(&result_count_BussHour);err != nil {
		return err
	}

	if result_count_BussHour < 1 {
		log.Println("error di sini brother: " + err.Error())
		return errors.New("failed")
	}

	query_insert_payments := `
		INSERT INTO payments (order_id, amount, payment_url, created_at)
		VALUES (:order_id, :amount, :payment_url, :created_at_payment);
	`

	paymets, err := tx.NamedExecContext(ctx, query_insert_payments, booking)
	if err != nil {
		return err
	}

	payment_id, err := paymets.LastInsertId()
	if err != nil {
		return err
	}

	booking.Payment_ID = int(payment_id)
	query_insert_booking := `INSERT INTO bookings (date, start_time, end_time, status_booking, user_id, service_id, payments_id, created_at)
		VALUES (:date, :start_time_booking, :end_time_booking, :status_booking, :user_id, :service_id, :payment_id, :created_at_booking);`
	
	if _, err := tx.NamedExecContext(ctx, query_insert_booking, booking);err != nil {
		return err
	}

	return tx.Commit()
}

func (d *database) GetPriceServices(ctx context.Context, serviceID int) (*domain.GetService, error) {
	var data domain.GetService
	query := `
		SELECT name,price FROM services WHERE id = ?
	`
	if err := d.db.GetContext(ctx, &data, query, serviceID); err != nil {
		return nil, err
	}
	return &data, nil
}
