package infrastructure

import (
	"github.com/Zyprush18/badmintonzz/internal/users/domain"
	"github.com/jmoiron/sqlx"
)

type UsersRepo interface {
	GetUsers() ([]domain.Users, error)
}


type repoUsers struct {
	db *sqlx.DB
}


func NewRepoUsers(d *sqlx.DB) UsersRepo {
	return &repoUsers{db: d}
}


func (u *repoUsers) GetUsers() ([]domain.Users, error) {
	var data []domain.Users

	if err:=  u.db.Select(&data, "SELECT * FROM users");err != nil {
		return nil, err
	}
	u.db.Close()

	return data, nil
}