package main

import (
	"github.com/bmizerany/pat"
	"github.com/princetomar27/basic_web_application/pkg/config"
	"github.com/princetomar27/basic_web_application/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	// 1. Create a multiplexer, which is a HTTP handler
	mux := pat.New()

	// 2. Set up your routes here
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
