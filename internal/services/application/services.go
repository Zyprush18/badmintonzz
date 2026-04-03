package application

import (
	"github.com/Zyprush18/badmintonzz/internal/services/application/queries"
	"github.com/Zyprush18/badmintonzz/internal/services/infrastructure"
)


type ApplicationServices interface {
	QueriesServices() queries.QueriesServices
}


type repoServices struct {
	repo infrastructure.RepoServices
}

func NewApplicationServices(r infrastructure.RepoServices) ApplicationServices {
	return &repoServices{
		repo: r,
	}
}

func (r *repoServices) QueriesServices() queries.QueriesServices {
	return queries.NewQueriesServices(r.repo)
}