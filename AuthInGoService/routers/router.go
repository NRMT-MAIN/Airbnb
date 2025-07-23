package routers

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router) 
}

func SetupRouter(UserRouter Router) *chi.Mux{
	chiRouter := chi.NewRouter()

	chiRouter.Use(middlewares.RateLimiter)
	chiRouter.Get("/ping" , controllers.PingHandeler)

	UserRouter.Register(chiRouter)

	return chiRouter
}