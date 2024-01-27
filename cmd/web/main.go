package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/handlers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

// Creates instances for both the application's system-wide config and Session manager
var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := setupComponents()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting the server on port %v...\n", port)

	server := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func setupComponents() error {
	// Enable value storing on the Session type
	gob.Register(models.Reservation{})

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
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.UseCache = false
	app.TemplateCache = templateCache

	// Initialize the handlers
	repo := handlers.NewRepo(&app)
	render.NewTemplates(&app)
	handlers.NewHandlers(repo)

	return nil
}
