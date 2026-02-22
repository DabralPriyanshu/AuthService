package main

import (
	"Auth/app"
	config "Auth/config/env"
)

func main() {
	config.Load()
	config := app.NewConfig()
	app := app.NewApplication(*config)
	app.Run()
}
