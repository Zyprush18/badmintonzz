package commands

import (
	"context"
	"time"

	"github.com/Zyprush18/badmintonzz/internal/shared/encrypt"
	"github.com/Zyprush18/badmintonzz/internal/users/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/users/interfaces/request"
)

type ServiceUsers interface {
	CreateUsers(ctx context.Context, user *request.UserRequest) error
	UpdateUsers(ctx context.Context, id int, user *request.UserRequest) error
	DeleteUsers(ctx context.Context, id int) error
}

type repoUsers struct {	
	repo infrastructure.UsersRepo
}

func GetCommandsUsers(r infrastructure.UsersRepo) ServiceUsers {
	return &repoUsers{repo: r}
}


func (r *repoUsers) CreateUsers(ctx context.Context, user *request.UserRequest) error {
	hashedPassword, err := encrypt.HashingPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return r.repo.CreateUser(ctx, user)
}


func (r *repoUsers) UpdateUsers(ctx context.Context, id int, user *request.UserRequest) error {
	hashedPassword, err := encrypt.HashingPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	request_data := map[string]interface{}{
		"id": id,
		"username": user.Username,
		"email": user.Email,
		"password": user.Password,
		"no_hp": user.Phone,
		"updated_at": time.Now(),
	}
	return r.repo.UpdateUser(ctx, request_data)
}


func (r *repoUsers) DeleteUsers(ctx context.Context, id int) error {
	return r.repo.DeleteUser(ctx, id)
}