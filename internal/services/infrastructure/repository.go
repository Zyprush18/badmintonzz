package infrastructure

import (
	"github.com/Zyprush18/badmintonzz/internal/services/domain"
	"github.com/jmoiron/sqlx"
)


type RepoServices interface {
	GetServices() ([]domain.Services, error)
}

type database struct {
	db *sqlx.DB
}

func NewRepoServices(d *sqlx.DB) RepoServices {
	return &database{
		db: d,
	}
}


func (d *database) GetServices() ([]domain.Services, error) {
	var services []domain.Services
	err := d.db.Select(&services, "SELECT * FROM services")
	if err != nil {
		return nil, err
	}
	return services, nil
}