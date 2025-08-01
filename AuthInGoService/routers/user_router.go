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
	r.With(middlewares.JWTMiddleware , middlewares.RequireAnyRoles("user" , "admin")).Get("/profile" , ur.userController.GetUserById)
	r.With(middlewares.ValidateRequestBody[dtos.CreateUserRequest]).Post("/signup" , ur.userController.Create)
	r.With(middlewares.ValidateRequestBody[dtos.LoginUserRequest]).Post("/login" , ur.userController.LoginUser)
}