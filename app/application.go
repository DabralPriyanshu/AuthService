package app

import (
	dbConfig "Auth/config/db"
	config "Auth/config/env"
	"Auth/controllers"
	db "Auth/db/repositories"
	"Auth/router"
	"Auth/services"
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
	Store  db.Storage
}

// constructor for config
func NewConfig() *Config {
	port :=
		config.GetString("PORT", ":8080")
	return &Config{Addr: port}

}

// constructor for Application
func NewApplication(config Config) *Application {
	return &Application{Config: config, Store: *db.NewStorage()}

}
func (app *Application) Run() error {
	dbInstance, err := dbConfig.SetupDB()
	if err != nil {
		fmt.Println("Error setting db", err)
	}
	ur := db.NewUserRepository(dbInstance)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(uRouter), //will return  chi router reference
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on", app.Config.Addr)
	return server.ListenAndServe()
}
