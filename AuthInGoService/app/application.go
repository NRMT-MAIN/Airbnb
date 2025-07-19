package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/routers"
	"fmt"
	"net/http"
	"time"
)

//Config hold the configuration of the Server
type Config struct {
	Addr string //PORT
}

type Application struct {
	Config Config
}

//Constructor
func NewConfig() Config {
	port := config.GetString("PORT" , ":8080")

	return Config{
		Addr: port,
	}
}

func NewApplication(cfg Config) Application {
	return Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: routers.SetupRouter(),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server is running on port :" , app.Config.Addr) ; 

	return server.ListenAndServe() ; 
}