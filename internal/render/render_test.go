package render

import (
	"net/http"
	"testing"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var templateData models.TemplateData

	response, err := getSession()
	if err != nil {
		t.Error(err)
	}

	testSession.Put(response.Context(), "flash", "Dummy flash message...")
	result := AddDefaultData(&templateData, response)

	if result == nil {
		t.Errorf("Expected %v, got %v", templateData, result)
	}

}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	templateCache, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	appConfig.TemplateCache = templateCache

	testRequest, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var testResponseWriter TestResponseWriter

	err = RenderTemplate(&testResponseWriter, testRequest, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Errorf("Error writing template to browser: %v", err)
	}

	err = RenderTemplate(&testResponseWriter, testRequest, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Errorf("Non-existent template should return error")
	}
}

func getSession() (*http.Request, error) {
	response, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	context := response.Context()
	context, _ = testSession.Load(context, response.Header.Get("X-Session"))

	// update the request with the new context
	response = response.WithContext(context)

	return response, nil
}

func TestNewTemplate(t *testing.T) {
	NewTemplates(appConfig)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
