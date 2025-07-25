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

func SetupRouter(UserRouter Router) *chi.Mux{
	chiRouter := chi.NewRouter()

	chiRouter.Use(middlewares.RateLimiter)
	chiRouter.HandleFunc("/fake-product-service/*" , utils.ReverseProxy("https://fakestoreapi.in" , "fake-product-service" ,"api"))
	chiRouter.Get("/ping" , controllers.PingHandeler)

	UserRouter.Register(chiRouter)

	return chiRouter
}
//https://fakestoreapi.in/api/products 