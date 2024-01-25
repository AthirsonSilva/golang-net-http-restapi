package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/pkg/models"
	"github.com/justinas/nosurf"
)

var appConfig *config.AppConfig

func NewTemplates(ac *config.AppConfig) {
	appConfig = ac
}

func AddDefaultData(templateData *models.TemplateData, r *http.Request) *models.TemplateData {
	templateData.CSRFToken = nosurf.Token(r)
	return templateData
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, templateFile string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template

	if appConfig.UseCache {
		templateCache = appConfig.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	templateCache = appConfig.TemplateCache
	template, ok := templateCache[templateFile]

	if !ok {
		log.Fatal("Template not found => ", templateFile)
	}

	buffer := new(bytes.Buffer)
	templateData = AddDefaultData(templateData, r)
	err := template.Execute(buffer, templateData)

	if err != nil {
		log.Fatal("Error executing template => ", err)
	}

	_, err = buffer.WriteTo(w)

	if err != nil {
		log.Println("Error writing template to browser => ", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)

		if err != nil {
			return templateCache, err
		}

		layoutMatches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return templateCache, err
		}

		if len(layoutMatches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = templateSet
	}

	return templateCache, nil
}
