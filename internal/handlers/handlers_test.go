package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type paramsData struct {
	key   string
	value string
}

type testStructure struct {
	name         string
	url          string
	method       string
	params       []paramsData
	expectedCode int
}

var testList = []testStructure{
	{"Index", "/", "GET", []paramsData{}, http.StatusOK},
	{"About", "/about", "GET", []paramsData{}, http.StatusOK},
	{"Contact", "/contact", "GET", []paramsData{}, http.StatusOK},
	{"General Quartes", "/generals-quarters", "GET", []paramsData{}, http.StatusOK},
	{"Major Suites", "/majors-suite", "GET", []paramsData{}, http.StatusOK},
	{"Search Availability", "/search-availability", "GET", []paramsData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getTestRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	for _, test := range testList {
		if test.method == "GET" {
			response, err := testServer.Client().Get(testServer.URL + test.url)
			if err != nil {
				t.Error(err)
				t.Fatal(err)
			}

			if response.StatusCode != test.expectedCode {
				t.Errorf("Expected status code to be %d, but got %d for the test %s",
					test.expectedCode, response.StatusCode, test.name)
			}
		}
	}	
}
