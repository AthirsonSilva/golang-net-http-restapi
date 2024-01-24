package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello world!"))
}

func About(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("About page")))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Page not found"))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var f, err = calculateDivide(10, 2)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("10 / 2 = %v", f)))
}

func calculateDivide(a, b int) (int, error) {
	if b == 0 || a == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/not-found", NotFound)
	http.HandleFunc("/divide", Divide)

	log.Printf("Starting the server on port %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
