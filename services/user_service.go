package services

import (
	config "Auth/config/env"
	db "Auth/db/repositories"
	"Auth/dto"
	"Auth/models"
	"Auth/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById(id string) (*models.User, error)
	CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error)
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
}

type UserServiceImp struct {
	userRepository db.UserRepository //it is depending on interface rather than the actual class (loose coupling )
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImp{userRepository: _userRepository}

}
func (u *UserServiceImp) GetUserById(id string) (*models.User, error) {
	fmt.Println("Fetching user in UserService")
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImp) CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error) {
	fmt.Println("Creating user in UserService")

	// Step 1. Hash the password using utils.HashPassword
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return nil, err
	}

	// Step 2. Call the repository to create the user
	user, err := u.userRepository.Create(payload.Username, payload.Email, hashedPassword)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return nil, err
	}

	// Step 3. Return the created user
	return user, nil
}

func (u *UserServiceImp) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	fmt.Println("Logging user in service")
	user, err := u.userRepository.GetByEmail(payload.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		fmt.Println("No user found with given email")
	}
	isPasswordValid := utils.ComparePassword(payload.Password, user.Password)
	if !isPasswordValid {
		fmt.Println("Wrong Password")
		return "", err

	}
	jwtPayload := jwt.MapClaims{"email": user.Email, "id": user.Id}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)
	tokenString, err := token.SignedString([]byte(config.GetString("JWT_SECRET", "TOKEN")))
	if err != nil {
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil
}
