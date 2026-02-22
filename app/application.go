package app

import (
	config "Auth/config/env"
	"Auth/router"
	"fmt"
	"net/http"
	"time"
)

// holds the configuration of server
type Config struct {
	Addr string
}

// it contains server details
type Application struct {
	Config Config
}

// constructor for config
func NewConfig() *Config {
	port :=
		config.GetString("PORT", ":8080")
	return &Config{Addr: port}

}

// constructor for Application
func NewApplication(config Config) *Application {
	return &Application{Config: config}

}
func (app *Application) Run() error {
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(),  //will return  chi router reference
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on ", app.Config.Addr)
	return server.ListenAndServe()
}
