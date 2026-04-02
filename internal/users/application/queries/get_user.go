package queries

import (
	"context"

	"github.com/Zyprush18/badmintonzz/internal/users/domain"
	"github.com/Zyprush18/badmintonzz/internal/users/infrastructure"
	"github.com/Zyprush18/badmintonzz/internal/users/interfaces/response"
)

type ServiceUsers interface {
	GetUsers(ctx context.Context) ([]response.UserResponse, error)
	GetUser(ctx context.Context, id int) (*response.UserResponse, error)
}

type userServices struct {
	repo infrastructure.UsersRepo
}


func NewServiceUsers(r infrastructure.UsersRepo) ServiceUsers {
	return &userServices{repo: r}
}

func (r *userServices) GetUsers(ctx context.Context) ([]response.UserResponse, error) {
	var data []response.UserResponse
	data_user, err := r.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	for _, user := range data_user {
		data = append(data, r.ConvertToResponse(&user))
	}
	return data, nil
}


func (r *userServices) GetUser(ctx context.Context, id int) (*response.UserResponse, error) {
	user, err := r.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	data := r.ConvertToResponse(user)

	return &data, nil
}


func (r *userServices) ConvertToResponse(user *domain.Users) response.UserResponse {
	return response.UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		No_Hp:      user.No_Phone,
		Created_at: user.Create_at,
		Updated_at: user.Updated_at,
	}
}