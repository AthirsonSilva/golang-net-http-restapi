package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/handlers"
	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.UseCache = false
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = templateCache

	repo := handlers.NewRepo(&app)
	render.NewTemplates(&app)
	handlers.NewHandlers(repo)

	log.Printf("Starting the server on port %v...\n", port)

	server := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
