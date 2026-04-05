package infrastructure

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/schedule/domain"
	"github.com/Zyprush18/badmintonzz/internal/schedule/interfaces/request"
	"github.com/jmoiron/sqlx"
)


type RepoSchedules interface {
	GetSchedules(ctx context.Context) ([]domain.ScheduleServices, error)
	GetSchedule(ctx context.Context, id int) (*domain.ScheduleServices, error)
	CreateSchedule(ctx context.Context, schedule *request.ScheduleRequest) error
	UpdateSchedule(ctx context.Context, schedule map[string]interface{}) error
	DeleteSchedule(ctx context.Context, id int) error
}


type database struct {
	db *sqlx.DB
}

func NewRepoSchedule(d *sqlx.DB) RepoSchedules {
	return &database{
		db: d,
	}
}

func (d *database) GetSchedules(ctx context.Context) ([]domain.ScheduleServices, error) {
	var schedules []domain.ScheduleServices
	query := `
	SELECT
		sch.id as schedule_id,
		sch.date as date,
		sch.time as time,
		sch.duration as duration,
		sch.created_at as created_at_scdl,
		sch.updated_at as updated_at_scdl,
		sch.service_id as service_id,
		svc.name as name,
		svc.price as price,
		svc.created_at as created_at_svc,
		svc.updated_at as updated_at_svc
	FROM schedules sch
	LEFT JOIN services svc ON sch.service_id = svc.id`	


	if err:= d.db.SelectContext(ctx, &schedules, query);err != nil {
		return nil, err
	}

	return schedules, nil
}


func (d *database) GetSchedule(ctx context.Context, id int) (*domain.ScheduleServices, error) {
	var schedule domain.ScheduleServices
	query := `
	SELECT
		sch.id as schedule_id,
		sch.date as date,
		sch.time as time,
		sch.duration as duration,
		sch.created_at as created_at_scdl,
		sch.updated_at as updated_at_scdl,
		sch.service_id as service_id,
		svc.name as name,
		svc.price as price,
		svc.created_at as created_at_svc,
		svc.updated_at as updated_at_svc
	FROM schedules sch
	LEFT JOIN services svc ON sch.service_id = svc.id
	WHERE sch.id = ?`

	if err := d.db.GetContext(ctx, &schedule, query, id); err != nil {
		return nil, err
	}

	return &schedule, nil
}


func (d *database) CreateSchedule(ctx context.Context, schedule *request.ScheduleRequest) error {
	query := `
		INSERT INTO schedules (date, time, duration, service_id)
		VALUES (:date, :time, :duration, :service_id)
	`
	_, err := d.db.NamedExecContext(ctx, query, schedule)
	return err
}


func (d *database) UpdateSchedule(ctx context.Context, schedule map[string]interface{}) error {
	query := `
		UPDATE schedules
		SET date = :date, time = :time, duration = :duration, service_id = :service_id
		WHERE id = :id
	`
	_, err := d.db.NamedExecContext(ctx, query, schedule)
	return err
}


func (d *database) DeleteSchedule(ctx context.Context, id int) error {
	query := `
		DELETE FROM schedules
		WHERE id = ?
	`
	_, err := d.db.ExecContext(ctx, query, id)
	return err
}