package main

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/justinas/nosurf"
)

// WriteToConsole logs the request data to the console
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("Handler called for method => %s", req.Method)
		log.Printf("Handler called for protocol => %s", req.Proto)
		log.Printf("Handler called for URL => %s%s", req.Host, req.URL.Path)

		headers := []string{"Content-Type", "Cookie", "Accept", "User-Agent"}
		for _, h := range headers {
			headerValue := req.Header.Get(h)
			if headerValue == "" {
				log.Printf("Handler called for header => %s: <empty>", h)
			} else {
				log.Printf("Handler called for header => %s: %s", h, headerValue)
			}
		}

		next.ServeHTTP(res, req)
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

// VerifyUserAuthentication verifies if the user is logged in
func VerifyUserAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if !helpers.IsAuthenticated(req) {
			log.Println("There is no user currently logged in")
			session.Put(req.Context(), "error", "Log in first!")
			http.Redirect(res, req, "/", http.StatusSeeOther)
			return
		}

		log.Println("The user is logged in!")
		next.ServeHTTP(res, req)
	})
}
