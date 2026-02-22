package controllers

import (
	"Auth/services"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {

	return &UserController{UserService: _userService}
}

func (uc *UserController) GetByUserId(w http.ResponseWriter, r *http.Request) {
	uc.UserService.GetByUserId()
	w.Write([]byte("User get profile endpoint hit"))

}
