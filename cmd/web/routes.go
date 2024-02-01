package main

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Creates a new Chi router and routes all available endpoints
func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(NoSurf)
	router.Use(SessionLoad)
	router.Use(WriteToConsole)

	router.Get("/", usecases.Repo.Home)
	router.Get("/about", usecases.Repo.About)
	router.Get("/contact", usecases.Repo.Contact)
	router.Get("/generals-quarters", usecases.Repo.Generals)
	router.Get("/majors-suite", usecases.Repo.Majors)

	router.Get("/reservation-summary", usecases.Repo.ReservationSummary)
	router.Get("/make-reservation", usecases.Repo.MakeReservation)
	router.Post("/make-reservation", usecases.Repo.PostReservation)

	router.Get("/search-availability", usecases.Repo.Availability)
	router.Post("/search-availability-json", usecases.Repo.PostAvailabilityJSON)
	router.Post("/search-availability", usecases.Repo.PostAvailability)
	router.Get("/choose-room/{id}", usecases.Repo.ChooseRoom)

	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}
