package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/princetomar27/basic_web_application/pkg/config"
	"github.com/princetomar27/basic_web_application/pkg/handlers"
	"net/http"
)

func chiRoutes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)

	mux.Use(NoSurve)
	mux.Use(SessionLoad)

	// Register the routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
