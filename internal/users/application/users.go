package application

import (
	"github.com/Zyprush18/badmintonzz/internal/users/application/commands"
	"github.com/Zyprush18/badmintonzz/internal/users/application/queries"
	"github.com/Zyprush18/badmintonzz/internal/users/infrastructure"
)

type ServicesUsers interface {
	QueriesUsers() queries.ServiceUsers
	CommandsUsers() commands.ServiceUsers
}

type RepoUser struct {
	Repo infrastructure.UsersRepo
}

func NewServiceUsers(r infrastructure.UsersRepo) ServicesUsers {
	return &RepoUser{Repo: r}
}

func (r *RepoUser) QueriesUsers() queries.ServiceUsers {
	return queries.GetQueriesUsers(r.Repo)
}

func (r *RepoUser) CommandsUsers() commands.ServiceUsers {
	return commands.GetCommandsUsers(r.Repo)
}