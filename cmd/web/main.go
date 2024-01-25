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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Starting the server on port %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
