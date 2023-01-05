package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sanyog2491/bookings/package/config"
	"github.com/sanyog2491/bookings/package/handlers"
)

func routes(app *config.Appconfig) http.Handler {
	//multiplexar which is an http handler using PAT for routing
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home3))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About2))
	//return mux

	//multiplexar which is an http handler using chi for routing

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(Nosurf)
	mux.Use(Sessionload)

	mux.Get("/", handlers.Repo.Home3)
	mux.Get("/about", handlers.Repo.About2)

	return mux
}
