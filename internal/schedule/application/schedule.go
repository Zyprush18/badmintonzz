package application

import (
	"github.com/Zyprush18/badmintonzz/internal/schedule/application/commands"
	"github.com/Zyprush18/badmintonzz/internal/schedule/application/queries"
	"github.com/Zyprush18/badmintonzz/internal/schedule/infrastructure"
)


type ScheduleApp interface {
	QueriesSchedules() queries.QueriesSchedules
	CommandsSchedules() commands.CommandsSchedule
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


func (r *repoSchedule) CommandsSchedules() commands.CommandsSchedule {
	return commands.NewCommandSchedule(r.repo)
}