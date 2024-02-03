package usecases

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}

// Repo the repository used by the handlers
func getTestRoutes() http.Handler {
	gob.Register(models.User{})
	gob.Register(models.Reservation{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(models.RoomRestriction{})

	// Change to true when in production
	app.InProduction = false

	// Initialize the session manager
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// Initialize the template cache
	templateCache, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.UseCache = true
	app.TemplateCache = templateCache

	// Initialize the handlers
	repo := NewRepo(&app, nil)
	render.NewRenderer(&app)
	NewHandlers(repo)

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(SessionLoad)
	router.Use(WriteToConsole)

	router.Get("/", Repo.Home)
	router.Get("/about", Repo.About)
	router.Get("/contact", Repo.Contact)
	router.Get("/generals-quarters", Repo.Generals)
	router.Get("/majors-suite", Repo.Majors)

	router.Get("/reservation-summary", Repo.ReservationSummary)
	router.Get("/make-reservation", Repo.MakeReservation)
	router.Post("/make-reservation", Repo.PostReservation)

	router.Get("/search-availability", Repo.Availability)
	router.Post("/search-availability-json", Repo.PostAvailabilityJSON)
	router.Post("/search-availability", Repo.PostAvailability)

	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return router
}

// NoSurf adds CSRF protection to all POST reqs
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session on every req
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// WriteToConsole logs the req data to the console
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handler called for method => %s", r.Method)
		log.Printf("Handler called for protocol => %s", r.Proto)
		log.Printf("Handler called for URL => %s%s", r.Host, r.URL.Path)

		headers := []string{"Content-Type", "Content-Length", "Accept-Encoding", "Cookie", "Accept", "User-Agent"}
		for _, h := range headers {
			headerValue := r.Header.Get(h)
			if headerValue == "" {
				log.Printf("Handler called for header => %s: <empty>", h)
			} else {
				log.Printf("Handler called for header => %s: %s", h, headerValue)
			}
		}

		next.ServeHTTP(w, r)
	})
}

// CreateTestTemplateCache creates a template cache as a map
func CreateTestTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))

	if err != nil {
		return templateCache, err
	}

	// Loop through all files ending with page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return templateCache, err
		}

		// Look for any layout files
		layoutMatches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))

		if err != nil {
			return templateCache, err
		}

		// If there are any layout files
		if len(layoutMatches) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return templateCache, err
			}
		}

		// Add template set to the cache
		templateCache[name] = templateSet
	}

	return templateCache, nil
}
