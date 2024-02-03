package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/alexedwards/scs/v2"
)

var testSession *scs.SessionManager
var testAppConfig config.AppConfig

func TestMain(m *testing.M) {
	// Enable value storing on the Session type
	gob.Register(models.User{})
	gob.Register(models.Reservation{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(models.RoomRestriction{})

	// Change to true when in production
	testAppConfig.InProduction = false

	// Initialize the session manager
	testSession = scs.New()
	testSession.Lifetime = 24 * time.Hour
	testSession.Cookie.Persist = true
	testSession.Cookie.SameSite = http.SameSiteLaxMode
	testSession.Cookie.Secure = testAppConfig.InProduction
	testAppConfig.Session = testSession

	// Set the app config to the test config
	TestAppConfig = &testAppConfig

	os.Exit(m.Run())
}

type TestResponseWriter struct{}

func (t *TestResponseWriter) Header() http.Header {
	var header http.Header
	return header
}

func (t *TestResponseWriter) WriteHeader(statusCode int) {
}

func (t *TestResponseWriter) Write(data []byte) (int, error) {
	length := len(data)
	if length == 0 {
		return 0, nil
	}
	return length, nil
}
