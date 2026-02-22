package router

import (
	"Auth/controllers"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	router := chi.NewRouter()
	//start defining routes
	router.Get("/ping", controllers.PingHandler)
	return router

}
