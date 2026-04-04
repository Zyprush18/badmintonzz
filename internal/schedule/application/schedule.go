package application

import (
	"github.com/Zyprush18/badmintonzz/internal/schedule/application/queries"
	"github.com/Zyprush18/badmintonzz/internal/schedule/infrastructure"
)


type ScheduleApp interface {
	QueriesSchedules() queries.QueriesSchedules
}

type repoSchedule struct {
	repo infrastructure.RepoSchedules
}

func NewApplicationSchedules(r infrastructure.RepoSchedules) ScheduleApp {
	return &repoSchedule{
		repo: r,
	}
}


func (r *repoSchedule) QueriesSchedules() queries.QueriesSchedules {
	return queries.NewQueriesSchedule(r.repo)
}