package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

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

// IsAuthenticated checks if the user is authenticated
func IsAuthenticated(responseWriter http.ResponseWriter, request *http.Request) bool {
	if app.Session.Exists(request.Context(), "user_id") {
		return true
	}
	http.Redirect(responseWriter, request, "/user/login", http.StatusSeeOther)
	return false
}

// ConvertDateFromString converts a string to a time.Time
func ConvertDateFromString(date string, responseWriter http.ResponseWriter) time.Time {
	layout := "2006-01-02"
	endDate, err := time.Parse(layout, date)
	if err != nil {
		ServerError(responseWriter, err)
	}
	return endDate
}

// JsonResponse returns a JSON response with passed HTTP status code
func JsonResponse(responseWriter http.ResponseWriter, status int, data interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(status)
	response, err := json.MarshalIndent(data, "", "     ")
	if err != nil {
		ServerError(responseWriter, err)
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(status)
	responseWriter.Write(response)
}
