package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
	dbConfig "AuthInGo/config/db"
)

func main(){
	config.Load() 
	dbConfig.SetupDB()

	cfg := app.NewConfig()

	app := app.NewApplication(cfg)
	app.Run() ; 
}