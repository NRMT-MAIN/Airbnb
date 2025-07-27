package main

import (
	"reviewservice/app"
)

func main() {
	cfg := app.NewConfig()

	app := app.NewApplication(cfg)

	app.Run() ; 
}