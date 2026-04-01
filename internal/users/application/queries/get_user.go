package queries

import (
	"github.com/Zyprush18/badmintonzz/internal/users/domain"
	"github.com/Zyprush18/badmintonzz/internal/users/infrastructure"
)

type ServiceUsers interface {
	GetUsers() ([]domain.Users, error)
}

type userServices struct {
	repo infrastructure.UsersRepo
}


func NewServiceUsers(r infrastructure.UsersRepo) ServiceUsers {
	return &userServices{repo: r}
}

func (r *userServices) GetUsers() ([]domain.Users, error) {
	return r.repo.GetUsers()
}