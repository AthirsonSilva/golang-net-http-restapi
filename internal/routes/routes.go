package routes

import (
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/middlewares"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Creates a new router to distribute all available endpoints
func Routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middlewares.NoSurf)
	router.Use(middlewares.SessionLoad)
	router.Use(middlewares.WriteToConsole)

	router.Get("/", usecases.Repo.Home)
	router.Get("/about", usecases.Repo.About)

	router.Route("/reservations", func(router chi.Router) {
		router.Use(middlewares.VerifyUserAuthentication)
		router.Get("/reservation-summary", usecases.Repo.ReservationSummary)
		router.Get("/make-reservation", usecases.Repo.MakeReservation)
		router.Post("/make-reservation", usecases.Repo.PostReservation)
	})

	router.Get("/search-availability", usecases.Repo.Availability)
	router.Post("/search-availability-json", usecases.Repo.PostAvailabilityJSON)
	router.Post("/search-availability", usecases.Repo.PostAvailability)

	router.Get("/choose-room/{id}", usecases.Repo.ChooseRoom)
	router.Get("/find-availability-by-room/{id}", usecases.Repo.FindAvailabilityByRoom)

	router.Get("/login", usecases.Repo.LoginPage)
	router.Post("/login", usecases.Repo.Login)
	router.Post("/register", usecases.Repo.Register)
	router.Get("/logout", usecases.Repo.Logout)

	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	router.Route("/admin", func(router chi.Router) {
		// router.Use(middlewares.VerifyUserAuthentication)
		router.Get("/dashboard", usecases.Repo.AdminDashboard)
		router.Get("/reservations/new", usecases.Repo.AdminAllNewReservations)
		router.Get("/reservations/all", usecases.Repo.AdminAllReservations)
		router.Get("/reservations/{id}", usecases.Repo.AdminShowSingleReservation)
		router.Get("/reservations/delete/{id}", usecases.Repo.AdminDeleteReservationByID)
		router.Post("/reservations/update", usecases.Repo.AdminUpdateReservation)
	})

	return router
}
