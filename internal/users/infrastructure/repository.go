package infrastructure

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/users/domain"
	"github.com/jmoiron/sqlx"
)

type UsersRepo interface {
	GetUsers(ctx context.Context) ([]domain.Users, error)
	GetUser(ctx context.Context, id int) (*domain.Users, error)
}


type repoUsers struct {
	db *sqlx.DB
}


func NewRepoUsers(d *sqlx.DB) UsersRepo {
	return &repoUsers{db: d}
}


func (u *repoUsers) GetUsers(ctx context.Context) ([]domain.Users, error) {
	var data []domain.Users
	if err:=  u.db.SelectContext(ctx, &data, "SELECT * FROM users");err != nil {
		return nil, err
	}
	
	return data, nil
}


func (u *repoUsers) GetUser(ctx context.Context, id int) (*domain.Users, error) {
	var user domain.Users
	if err := u.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = ?", id); err != nil {
		return nil, err
	}

	return &user, nil
}