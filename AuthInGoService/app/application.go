package app

import (
	config "AuthInGo/config/env"
	dbconfig "AuthInGo/config/db"
	"AuthInGo/controllers"
	db "AuthInGo/db/repositories"
	"AuthInGo/routers"
	"AuthInGo/services"
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
	Store db.Storage
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
		Store: *db.NewStorage(),
	}
}

func (app *Application) Run() error {
	repo , err := dbconfig.SetupDB()

	if err != nil {
		fmt.Println("Error setting up DB connection" , err)
		return err
	}


	ur := db.NewUserRepository(repo)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := routers.NewUserRouter(uc)


	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: routers.SetupRouter(uRouter),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Server is running on port :" , app.Config.Addr) ; 

	return server.ListenAndServe() ; 
}