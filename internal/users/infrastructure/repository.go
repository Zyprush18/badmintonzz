package infrastructure

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/shared/cntx"
	"github.com/Zyprush18/badmintonzz/internal/users/domain"
	"github.com/Zyprush18/badmintonzz/internal/users/interfaces/request"
	"github.com/jmoiron/sqlx"
)

type UsersRepo interface {
	GetUsers(ctx context.Context) ([]domain.Users, error)
	GetUser(ctx context.Context, id int) (*domain.Users, error)
	CreateUser(ctx context.Context, user *request.UserRequest) error
	UpdateUser(ctx context.Context, user map[string]interface{}) error
	DeleteUser(ctx context.Context, id int) error
}


type repoUsers struct {
	db *sqlx.DB
}


func NewRepoUsers(d *sqlx.DB) UsersRepo {
	return &repoUsers{db: d}
}


func (u *repoUsers) GetUsers(ctx context.Context) ([]domain.Users, error) {
	ctx, cancel := cntx.TimeOutContext(ctx, cntx.DbMediumTime)
	defer cancel()

	var data []domain.Users
	if err:=  u.db.SelectContext(ctx, &data, "SELECT * FROM users");err != nil {
		return nil, err
	}
	
	return data, nil
}


func (u *repoUsers) GetUser(ctx context.Context, id int) (*domain.Users, error) {
	ctx, cancel := cntx.TimeOutContext(ctx, cntx.DbShortTime)
	defer cancel()

	var user domain.Users
	if err := u.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = ?", id); err != nil {
		return nil, err
	}

	return &user, nil
}


func (u *repoUsers) CreateUser(ctx context.Context, user *request.UserRequest) error {
	ctx, cancel := cntx.TimeOutContext(ctx, cntx.DbMediumTime)
	defer cancel()

	_, err := u.db.NamedExecContext(ctx, "INSERT INTO users (username, email, password, no_hp) VALUES (:username, :email, :password, :no_hp)", user)
	return err
}


func (u *repoUsers) UpdateUser(ctx context.Context, user map[string]interface{}) error {
	ctx, cancel := cntx.TimeOutContext(ctx, cntx.DbMediumTime)
	defer cancel()

	_, err := u.db.NamedExecContext(ctx, "UPDATE users SET username = :username, email = :email, password = :password, no_hp = :no_hp, updated_at = :updated_at WHERE id = :id", user)
	return err
}

func (u *repoUsers) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := cntx.TimeOutContext(ctx, cntx.DbShortTime)
	defer cancel()

	_, err := u.db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
	return err
}