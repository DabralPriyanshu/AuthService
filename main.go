package main

import (
	"Auth/app"
)

func main() {
	config := app.NewConfig(":3003")
	app := app.NewApplication(*config)
	app.Run()
}
