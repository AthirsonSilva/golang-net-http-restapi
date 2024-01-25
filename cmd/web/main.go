package main

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Starting the server on port %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
