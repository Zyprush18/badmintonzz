package commands

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/shared/encrypt"
	"github.com/Zyprush18/badmintonzz/internal/users/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/users/interfaces/request"
)

type ServiceUsers interface {
	CreateUsers(ctx context.Context, user *request.UserRequest) error
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