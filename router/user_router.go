package router

import (
	"Auth/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController *controllers.UserController
}

func (ur *UserRouter) Register(r chi.Router) {
	r.Get("/profile", ur.UserController.GetByUserId)

}

func NewUserRouter(_userController *controllers.UserController) Router {

	return &UserRouter{UserController: _userController}
}
