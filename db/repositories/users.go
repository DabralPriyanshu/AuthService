package db

import (
	"fmt"
)

type UserRepository interface {
	Create() error
}

type UserRepositoryImp struct {
	// db *sql.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImp{}
}

func (u *UserRepositoryImp) Create() error {
	fmt.Println("Creating user in UserRepository")
	return nil
}
