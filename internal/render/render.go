package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
	"github.com/justinas/nosurf"
)

// Creates empty variable for the System-wide configuration typ
var (
	app             *config.AppConfig
	pathToTemplates = "./templates"
	functions       = template.FuncMap{}
)

// Creates a new instance of the Templates function
func NewRenderer(ac *config.AppConfig) {
	app = ac
}

// AddDefaultData adds data for all templates
func AddDefaultData(templateData *models.TemplateData, req *http.Request) *models.TemplateData {
	templateData.Flash = app.Session.PopString(req.Context(), "flash")
	templateData.Error = app.Session.PopString(req.Context(), "error")
	templateData.Warning = app.Session.PopString(req.Context(), "warning")
	templateData.CSRFToken = nosurf.Token(req)

	if app.Session.Exists(req.Context(), "user_id") {
		templateData.IsAuthenticated = true
	}

	return templateData
}

// RenderTemplate renders templates using html/template with caching
func RenderTemplate(
	res http.ResponseWriter,
	req *http.Request,
	templateFile string,
	templateData *models.TemplateData,
) error {
	var templateCache map[string]*template.Template

	// Check if the cache is being used
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	template, ok := templateCache[templateFile]

	if !ok {
		errorMsg := fmt.Sprintf("The template %s does not exist", templateFile)
		log.Println(errorMsg)
		return errors.New(errorMsg)
	}

	buffer := new(bytes.Buffer)
	templateData = AddDefaultData(templateData, req)
	_ = template.Execute(buffer, templateData)
	_, err := buffer.WriteTo(res)
	if err != nil {
		log.Println("Error writing template to browser => ", err)
		return err
	}

	return nil
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return templateCache, err
	}

	// Loop through all files ending with page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		// Look for any layout files
		layoutMatches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return templateCache, err
		}

		// If there are any layout files
		if len(layoutMatches) > 0 {
			templateSet, err = templateSet.ParseGlob(
				fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates),
			)
			if err != nil {
				return templateCache, err
			}
		}

		// Add template set to the cache
		templateCache[name] = templateSet
	}

	return templateCache, nil
}
