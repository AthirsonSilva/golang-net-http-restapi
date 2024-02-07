package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
)

var app *config.AppConfig

// NewHelpers sets the app config
func NewHelpers(appConfig *config.AppConfig) {
	app = appConfig
}

// ClientError sends a specific status code and error message
func ClientError(res http.ResponseWriter, statusCode int) {
	http.Error(res, http.StatusText(statusCode), statusCode)
	app.InfoLog.Println("Client error with status code of", statusCode)
}

// ServerError sends a more detailed message containing the error and the stack trace
func ServerError(res http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(
		res,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

// IsAuthenticated checks if the user is authenticated
func IsAuthenticated(req *http.Request) bool {
	if app.Session.Exists(req.Context(), "user_id") {
		return true
	}
	return false
}

// ConvertDateFromString converts a string to a time.Time
func ConvertDateFromString(date string, res http.ResponseWriter) time.Time {
	layout := "2006-01-02"
	endDate, err := time.Parse(layout, date)
	if err != nil {
		ServerError(res, err)
	}
	return endDate
}

// JsonResponse returns a JSON response with passed HTTP status code
func JsonResponse(res http.ResponseWriter, status int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	response, err := json.MarshalIndent(data, "", "     ")
	if err != nil {
		ServerError(res, err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(response)
}

// GetPathVariableFromRequest extracts the last path variable from request and returns it
func GetPathVariableFromRequest(req *http.Request) string {
	path := strings.Split(req.URL.Path, "/")
	lastIndex := len(path) - 1
	pathVar := path[lastIndex]
	return pathVar
}
