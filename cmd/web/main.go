package main

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/handlers"
	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig
	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

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
