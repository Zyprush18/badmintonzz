package queries

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/schedule/interfaces/response"
	"github.com/Zyprush18/badmintonzz/internal/schedule/infrastructure"
	svc_domain "github.com/Zyprush18/badmintonzz/internal/services/domain"
)

type QueriesSchedules interface {
	GetSchedules(ctx context.Context) ([]response.Schedules, error)
	GetSchedule(ctx context.Context, id int) (*response.Schedules, error)
}


type repoSchedule struct {
	repo infrastructure.RepoSchedules
}


func NewQueriesSchedule(r infrastructure.RepoSchedules) QueriesSchedules {
	return &repoSchedule{
		repo: r,
	}
}

func (r *repoSchedule) GetSchedules(ctx context.Context) ([]response.Schedules, error) {
	data, err := r.repo.GetSchedules(ctx)
	if err != nil {
		return nil, err
	}


	var schedules []response.Schedules
	for _, v := range data {
		schedules = append(schedules, response.Schedules{
			ID: v.IdSchedule,
			Date: v.Date,
			Time: v.Time,
			Duration: v.Duration,
			Service_id: v.IdServices,
			Services: svc_domain.Services{
				ID: v.IdServices,
				Name: v.Name,
				Price: v.Price,
				Created_at: v.CreatedAt_Svc,
				Updated_at: v.UpdatedAt_Svc,
			},
			Created_at: v.CreatedAt_Scdl,
			Updated_at: v.UpdatedAt_Scdl,

		})
	}

	return schedules, nil
}

func (r *repoSchedule) GetSchedule(ctx context.Context, id int) (*response.Schedules, error) {
	data, err := r.repo.GetSchedule(ctx, id)
	if err != nil {
		return nil, err
	}

	schedule := &response.Schedules{
		ID: data.IdSchedule,
		Date: data.Date,
		Time: data.Time,
		Duration: data.Duration,
		Service_id: data.IdServices,
		Services: svc_domain.Services{
				ID: data.IdServices,
				Name: data.Name,
				Price: data.Price,
				Created_at: data.CreatedAt_Svc,
				Updated_at: data.UpdatedAt_Svc,
			},
			Created_at: data.CreatedAt_Scdl,
			Updated_at: data.UpdatedAt_Scdl,
	}


	return schedule, nil
}