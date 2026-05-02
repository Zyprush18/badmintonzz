package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/Zyprush18/badmintonzz/internal/shared/encrypt"
	"github.com/Zyprush18/badmintonzz/internal/users/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/users/interfaces/request"
)

type ServiceUsers interface {
	AuthLogin(ctx context.Context, user *request.UserAuthLoginRequest) (string,error)
	AuthRegister(ctx context.Context, user *request.UserAuthRegisterRequest) error
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


func (r *repoUsers) AuthRegister(ctx context.Context, user *request.UserAuthRegisterRequest) error {
	hashedPassword, err := encrypt.HashingPassword(user.Password)
	if err != nil {
		return err
	}

	req := &request.UserRequest{
		Username: user.Username,
		Email: user.Email,
		Password: hashedPassword,
		Phone: user.Phone,
		Role: "user",
	}

	return r.repo.CreateUser(ctx, req)
}


func (r *repoUsers) AuthLogin(ctx context.Context, user *request.UserAuthLoginRequest) (string,error) {
	data, err := r.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	if !encrypt.CheckHashingPassword(data.Password, user.Password) {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := encrypt.GenerateJWToken(data.ID, data.Username, data.Role)
	if err != nil {
		return "", err
	}

	return token, nil
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