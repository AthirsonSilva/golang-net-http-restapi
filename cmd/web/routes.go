package main

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(NoSurf)
	router.Use(SessionLoad)
	router.Use(WriteToConsole)

	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	router.Get("/make-reservation", handlers.Repo.Reservation)
	router.Get("/search-availability", handlers.Repo.Availability)
	router.Get("/contact", handlers.Repo.Contact)
	router.Get("/generals-quarters", handlers.Repo.Generals)
	router.Get("/majors-suite", handlers.Repo.Majors)

	router.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	router.Post("/search-availability", handlers.Repo.PostAvailability)

	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}
