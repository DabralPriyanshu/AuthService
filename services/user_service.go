package services

import (
	db "Auth/db/repositories"

	"fmt"
)

type UserService interface {
	GetByUserId() error
}

type UserServiceImp struct {
	userRepository db.UserRepository //it is depending on interface rather than the actual class (loose coupling )
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImp{userRepository: _userRepository}

}
func (u *UserServiceImp) GetByUserId() error {
	u.userRepository.Create()
	fmt.Println("Fetching user in UserService")
	return nil
}
