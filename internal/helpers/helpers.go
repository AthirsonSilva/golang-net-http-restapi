package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
)

var app *config.AppConfig

// NewHelpers sets the app config
func NewHelpers(appConfig *config.AppConfig) {
	app = appConfig
}

// ClientError sends a specific status code and error message
func ClientError(responseWriter http.ResponseWriter, statusCode int) {
	http.Error(responseWriter, http.StatusText(statusCode), statusCode)
	app.InfoLog.Println("Client error with status code of", statusCode)
}

// ServerError sends a more detailed message containing the error and the stack trace
func ServerError(responseWriter http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// GetAppConfig returns the app config
func GetAppConfig() *config.AppConfig {
	return app
}
