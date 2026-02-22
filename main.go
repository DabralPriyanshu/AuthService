package main

import (
	"Auth/app"
)

func main() {
	config := app.Config{Addr: ":3003"}
	app := app.Application{Config: config}
	app.Run()
}
