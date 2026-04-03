package queries

import (
	"github.com/Zyprush18/badmintonzz/internal/services/domain"
	"github.com/Zyprush18/badmintonzz/internal/services/infrastructure"
)

type QueriesServices interface {
	GetServices() ([]domain.Services, error)
}


type repoServices struct {
	repo infrastructure.RepoServices
}

func NewQueriesServices(r infrastructure.RepoServices) QueriesServices {
	return &repoServices{
		repo: r,
	}
}


func (r *repoServices) GetServices() ([]domain.Services, error) {
	return r.repo.GetServices()
}