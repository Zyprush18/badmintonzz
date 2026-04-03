package commands

import (
	"context"
	"time"

	"github.com/Zyprush18/badmintonzz/internal/services/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/services/interfaces/request"
)

type CommandsService interface {
	CreateService(ctx context.Context, service *request.Services) error
	UpdateService(ctx context.Context, id int,service *request.Services) error
	DeleteService(ctx context.Context, id int) error
}

type repoServices struct {
	repo infrastructure.RepoServices
}


func NewCommandsService(r infrastructure.RepoServices) CommandsService {
	return &repoServices{
		repo: r,
	}
}


func (r *repoServices) CreateService(ctx context.Context, service *request.Services) error {
	return r.repo.CreateService(ctx, service)
}

func (r *repoServices) UpdateService(ctx context.Context, id int,service *request.Services) error {
	dataRequest := map[string]interface{}{
		"id": id,
		"name": service.Name,
		"price": service.Price,
		"updated_at": time.Now(),
	}
	return r.repo.UpdateService(ctx, dataRequest)
}


func (r *repoServices) DeleteService(ctx context.Context, id int) error {
	return r.repo.DeleteService(ctx, id)
}