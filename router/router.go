package router

import (
	"Auth/controllers"
	"Auth/middlewares"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(UserRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()
	// chiRouter.Use(middlewares.RequestLogger)
	chiRouter.Use(middleware.Logger)
	chiRouter.Use(middlewares.RateLimitMiddleware)
	//start defining routes
	chiRouter.Get("/ping", controllers.PingHandler)
	UserRouter.Register(chiRouter)
	return chiRouter

}
