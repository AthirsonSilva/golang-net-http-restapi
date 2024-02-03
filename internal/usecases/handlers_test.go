package usecases

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type reqData struct {
	key   string
	value string
}

type testStructure struct {
	name         string
	url          string
	method       string
	params       []reqData
	expectedCode int
}

var testList = []testStructure{
	{"Index", "/", "GET", []reqData{}, http.StatusOK},
	{"About", "/about", "GET", []reqData{}, http.StatusOK},
	{"Contact", "/contact", "GET", []reqData{}, http.StatusOK},
	{"General Quartes", "/generals-quarters", "GET", []reqData{}, http.StatusOK},
	{"Major Suites", "/majors-suite", "GET", []reqData{}, http.StatusOK},
	{"Search Availability", "/search-availability", "GET", []reqData{}, http.StatusOK},
	{"Post Availability", "/search-availability", "POST", []reqData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"Post Availability", "/search-availability-json", "POST", []reqData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"Make Reservation", "/make-reservation", "POST", []reqData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "jsmith@gmail.com"},
	}, http.StatusOK},
	// {"Reservation Summary", "/reservation-summary", "GET", []reqData{}, http.StatusSeeOther},
}

func TestHandlers(t *testing.T) {
	routes := getTestRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	for _, test := range testList {
		switch test.method {

		case "GET":
			response, err := testServer.Client().Get(testServer.URL + test.url)
			if err != nil {
				t.Error(err)
			}

			if response.StatusCode != test.expectedCode {
				t.Errorf("Expected status code to be %d, but got %d for the test %s",
					test.expectedCode, response.StatusCode, test.name)
			}

		case "POST":
			values := url.Values{}
			for _, data := range test.params {
				values.Add(data.key, data.value)
			}

			response, err := testServer.Client().PostForm(testServer.URL+test.url, values)
			if err != nil {
				t.Error(err)
			}

			if response.StatusCode != test.expectedCode {
				t.Errorf("Expected status code to be %d, but got %d for the test %s",
					test.expectedCode, response.StatusCode, test.name)
			}

		default:
			t.Errorf("The method %s for the test %s is not supported", test.method, test.name)
		}
	}
}
