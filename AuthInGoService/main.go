package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
	dbConfig "AuthInGo/config/db"
)

func main(){
	config.Load() 

	cfg := app.NewConfig()
	dbConfig.InitDB()

	app := app.NewApplication(cfg)
	app.Run() ; 
}