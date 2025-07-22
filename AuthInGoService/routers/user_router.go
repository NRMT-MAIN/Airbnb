package routers

import (
	"AuthInGo/controllers"
	"AuthInGo/dtos"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		userController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.Get("/profile" , ur.userController.GetUserById)
	r.Post("/signup" , middlewares.ValidateRequestBody[dtos.CreateUserRequest](ur.userController.Create))
	r.Post("/login" , middlewares.ValidateRequestBody[dtos.LoginUserRequest](ur.userController.LoginUser))
}