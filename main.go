package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Starting the application...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			log.Print("GET request")
		case "POST":
			log.Print("POST request")
		case "PUT":
			log.Print("PUT request")
		case "DELETE":
			log.Print("DELETE request")
		default:
			log.Print("Unknown request")
		}

		log.Print(r.Header, r.Body, r.URL.Path)
		w.Write([]byte("Hello world!"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
