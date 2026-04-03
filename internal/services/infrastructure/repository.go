package infrastructure

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/services/domain"
	"github.com/Zyprush18/badmintonzz/internal/services/interfaces/request"
	"github.com/jmoiron/sqlx"
)


type RepoServices interface {
	GetServices(ctx context.Context) ([]domain.Services, error)
	GetServiceByID(ctx context.Context, id int) (*domain.Services, error)
	CreateService(ctx context.Context, service *request.Services) error
	UpdateService(ctx context.Context, service map[string]interface{}) error
	DeleteService(ctx context.Context, id int) error
}

type database struct {
	db *sqlx.DB
}

func NewRepoServices(d *sqlx.DB) RepoServices {
	return &database{
		db: d,
	}
}


func (d *database) GetServices(ctx context.Context) ([]domain.Services, error) {
	var services []domain.Services
	err := d.db.SelectContext(ctx, &services, "SELECT * FROM services")
	if err != nil {
		return nil, err
	}
	return services, nil
}


func (d *database) GetServiceByID(ctx context.Context, id int) (*domain.Services, error) {
	var service domain.Services
	err := d.db.GetContext(ctx, &service, "SELECT * FROM services WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &service, nil
}


func (d *database) CreateService(ctx context.Context, service *request.Services) error {
	_, err := d.db.NamedExecContext(ctx, "INSERT INTO services (name, price) VALUES (:name, :price)", service)
	if err != nil {
		return err
	}
	return nil
}


func (d *database) UpdateService(ctx context.Context, service map[string]interface{}) error {
	_,err := d.db.NamedExecContext(ctx, "UPDATE services SET name = :name, price = :price, updated_at = :updated_at WHERE id = :id", service)
	if err != nil {
		return err
	}
	return nil
}


func (d *database) DeleteService(ctx context.Context, id int) error {
	_, err := d.db.ExecContext(ctx, "DELETE FROM services WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}