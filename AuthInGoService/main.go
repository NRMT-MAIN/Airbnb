package main

import (
	"AuthInGo/app"
)

func main(){
	cfg := app.Config{
		Addr : ":3000" , 
	}

	app := app.Application{
		Config: cfg,
	}

	app.Run() ; 
}