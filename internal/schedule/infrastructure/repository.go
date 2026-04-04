package infrastructure

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/schedule/domain"
	"github.com/jmoiron/sqlx"
)


type RepoSchedules interface {
	GetSchedules(ctx context.Context) ([]domain.ScheduleServices, error)
	GetSchedule(ctx context.Context, id int) (*domain.ScheduleServices, error)
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