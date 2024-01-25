package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole logs the request data to the console
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handler called for URL => %s", r.URL.Path)
		log.Printf("Handler called for method => %s", r.Method)
		log.Printf("Handler called for protocol => %s", r.Proto)
		log.Printf("Handler called for host => %s", r.Host)
		log.Printf("Handler called for remote address => %s", r.RemoteAddr)

		for name, values := range r.Header {
			for _, value := range values {
				log.Printf("Handler called for header => %s: %s", name, value)
			}
		}

		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection to all POST requests
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

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
