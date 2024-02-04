package middlewares

import (
	"log"
	"net/http"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/helpers"
	"github.com/justinas/nosurf"
)

// WriteToConsole logs the request data to the console
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Printf("Request method => %s", req.Method)
		log.Printf("Request protocol => %s", req.Proto)
		log.Printf("Request URL => %s%s", req.Host, req.URL.Path)

		headers := []string{"Content-Type", "Cookie", "Accept", "User-Agent"}
		for _, h := range headers {
			headerValue := req.Header.Get(h)
			if headerValue == "" {
				log.Printf("Request header => %s: <empty>", h)
			} else {
				log.Printf("Request header => %s: %s", h, headerValue)
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
		Secure:   config.App.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return config.Session.LoadAndSave(next)
}

// VerifyUserAuthentication verifies if the user is logged in
func VerifyUserAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if !helpers.IsAuthenticated(req) {
			log.Println("There is no user currently logged in")
			config.Session.Put(req.Context(), "error", "Log in first!")
			http.Redirect(res, req, "/", http.StatusSeeOther)
			return
		}

		log.Println("The user is logged in!")
		next.ServeHTTP(res, req)
	})
}
