package routers

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"
	"AuthInGo/utils"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router) 
}

func SetupRouter(RoleRouter Router , UserRouter Router) *chi.Mux{
	chiRouter := chi.NewRouter()

	chiRouter.Use(middlewares.RateLimiter)
	chiRouter.HandleFunc("/fake-product-service/*" , utils.ReverseProxy("https://fakestoreapi.in" , "fake-product-service" ,"api"))
	chiRouter.Get("/ping" , controllers.PingHandeler)

	UserRouter.Register(chiRouter)
	RoleRouter.Register(chiRouter)

	return chiRouter
}
