package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/database"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/render"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/usecases"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

// Creates instances for both the application's system-wide config and Session manager
var app config.AppConfig
var session *scs.SessionManager

func main() {
	db, err := setupComponents()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	log.Printf("Starting the server on port %v...\n", port)

	server := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func setupComponents() (*database.Database, error) {
	// Enable value storing on the Session type
	gob.Register(models.User{})
	gob.Register(models.Reservation{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(models.RoomRestriction{})

	// Change to true when in production
	app.InProduction = false

	// Initialize loggers
	app.InfoLog = log.New(os.Stdout, "INFO => ", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR => ", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize the session manager
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// connecting to database
	log.Println("Connecting to database...")
	db, err := database.ConnectSQL(fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		"localhost",
		"5432",
		"postgres",
		"root",
		"bookings",
	))

	if err != nil {
		log.Fatal("cannot connect to database")
		return nil, err
	}

	log.Println("Connected to database!")

	// Initialize the template cache
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.UseCache = false
	app.TemplateCache = templateCache

	// Initialize the handlers
	repo := usecases.NewRepo(&app, db)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)
	usecases.NewHandlers(repo)

	return db, nil
}
