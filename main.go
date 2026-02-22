package main

import (
	"Auth/app"
	dbConfig "Auth/config/db"
	config "Auth/config/env"
)

func main() {
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(*cfg)
	dbConfig.SetupDB()
	app.Run()
}
