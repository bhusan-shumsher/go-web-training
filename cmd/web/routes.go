package main

import (
	"net/http"

	"github.com/bhusan-shumsher/go-web-training/pkg/config"
	"github.com/bhusan-shumsher/go-web-training/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
