package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type requestData struct {
	key   string
	value string
}

type testStructure struct {
	name         string
	url          string
	method       string
	params       []requestData
	expectedCode int
}

var testList = []testStructure{
	{"Index", "/", "GET", []requestData{}, http.StatusOK},
	{"About", "/about", "GET", []requestData{}, http.StatusOK},
	{"Contact", "/contact", "GET", []requestData{}, http.StatusOK},
	{"General Quartes", "/generals-quarters", "GET", []requestData{}, http.StatusOK},
	{"Major Suites", "/majors-suite", "GET", []requestData{}, http.StatusOK},
	{"Search Availability", "/search-availability", "GET", []requestData{}, http.StatusOK},
	{"Post Availability", "/search-availability", "POST", []requestData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"Post Availability", "/search-availability-json", "POST", []requestData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"Make Reservation", "/make-reservation", "POST", []requestData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "jsmith@gmail.com"},
	}, http.StatusOK},
	// {"Reservation Summary", "/reservation-summary", "GET", []requestData{}, http.StatusSeeOther},
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
