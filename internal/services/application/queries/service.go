package queries

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/services/domain"
	"github.com/Zyprush18/badmintonzz/internal/services/infrastructure"
)

type QueriesServices interface {
	GetServices(ctx context.Context) ([]domain.Services, error)
	GetServiceByID(ctx context.Context, id int) (*domain.Services, error)
}


type repoServices struct {
	repo infrastructure.RepoServices
}

func NewQueriesServices(r infrastructure.RepoServices) QueriesServices {
	return &repoServices{
		repo: r,
	}
}


func (r *repoServices) GetServices(ctx context.Context) ([]domain.Services, error) {
	return r.repo.GetServices(ctx)
}

func (r *repoServices) GetServiceByID(ctx context.Context, id int) (*domain.Services, error) {
	return r.repo.GetServiceByID(ctx, id)
}