package controllers

import (
	"Auth/dto"
	"Auth/services"
	"Auth/utils"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {

	return &UserController{UserService: _userService}
}

func (uc *UserController) GetByUserId(w http.ResponseWriter, r *http.Request) {

	uc.UserService.GetUserById(12)
	w.Write([]byte("User get profile endpoint hit"))

}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	payload :=
		r.Context().Value("payload").(dto.CreateUserRequestDTO)

	fmt.Println("Payload received:", payload)

	user, err := uc.UserService.CreateUser(&payload)
	fmt.Println(err)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusCreated, "User created successfully", user)
	fmt.Println("User created successfully:", user)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	payload :=
		r.Context().Value("payload").(dto.LoginUserRequestDTO)

	fmt.Println("Payload received:", payload)
	jwtToken, err := uc.UserService.LoginUser(&payload)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to login user", err)
		return

	}
	utils.WriteSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)

}
