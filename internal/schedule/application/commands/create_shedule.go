package commands

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/schedule/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/schedule/interfaces/request"
)

type CommandsSchedule interface {
	CreateSchedule(ctx context.Context, schedule *request.ScheduleRequest) error
	UpdateSchedule(ctx context.Context, id int,schedule *request.ScheduleRequest) error
	DeleteSchedule(ctx context.Context, id int) error
}

type repoSchedule struct {
	repo infrastructure.RepoSchedules
}


func NewCommandSchedule(r infrastructure.RepoSchedules) CommandsSchedule {
	return &repoSchedule{repo: r}
}



func (r *repoSchedule) CreateSchedule(ctx context.Context, schedule *request.ScheduleRequest) error {
	return r.repo.CreateSchedule(ctx, schedule)
}


func (r *repoSchedule) UpdateSchedule(ctx context.Context, id int,schedule *request.ScheduleRequest) error  {
	data := map[string]interface{}{
		"id": id,
		"date": schedule.Date,
		"time": schedule.Time,
		"duration": schedule.Duration,
		"service_id": schedule.Service_id,
	}	
	return r.repo.UpdateSchedule(ctx, data)
}


func (r *repoSchedule) DeleteSchedule(ctx context.Context, id int) error  {
	return r.repo.DeleteSchedule(ctx, id)
}
