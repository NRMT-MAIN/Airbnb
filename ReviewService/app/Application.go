package app

import (
	"fmt"
	"net/http"
	config "reviewservice/config/db"
	"reviewservice/config/env"
	"reviewservice/controllers"
	db "reviewservice/db/repositories"
	"reviewservice/routers"
	"reviewservice/services"
	"time"
)

type Config struct {
	Addr string
	Store db.Storage
}

type Application struct {
	Config Config
}

func NewConfig() Config {
	port := env.GetString("PORT" , ":3000")
	return Config{
		Addr: port,
		Store: *db.NewStorage(),
	}
}

func NewApplication(cfg Config) Application {
	return Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {
	dbSetup , err := config.SetupDB()

	if err != nil {
		fmt.Println("Database not setup")
		return err 
	}

	rep := db.NewReviewRepository(dbSetup)
	rs := services.NewReviewService(rep)
	rc := controllers.NewReviewController(rs)
	rr := routers.NewReviewRouter(rc)


	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: routers.SetupRouter(rr) , 
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server is running on PORT :" , app.Config.Addr)

	return server.ListenAndServe() ; 
}