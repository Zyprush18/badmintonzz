package users

import "database/sql"

type UsersRepo interface {
	GetUsers() ([]Users, error)
}


type RepoUsers struct {
	db *sql.DB
}


func NewRepoUsers(d *sql.DB) *RepoUsers {
	return &RepoUsers{db: d}
}


func (r *RepoUsers) GetUsers() ([]Users, error) {
	// Implementation for fetching users from the database
	return nil, nil
}