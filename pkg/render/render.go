package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, templateFile string) {
	var parsedTemplate, _ = template.ParseFiles("./templates/"+templateFile, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Println("Error parsing template: ", err)
	}
}

var templateCache = make(map[string]*template.Template)

func NewRenderTemplate(w http.ResponseWriter, templateFile string) {
	var template *template.Template
	var err error

	// Check if the template is in cache
	_, inMap := templateCache[templateFile]

	if !inMap {
		log.Println("Creating template and adding to cache")
		err = createTemplateCache(templateFile)

		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Using cached template")
	}

	template = templateCache[templateFile]
	err = template.Execute(w, nil)

	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	var templates = []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// Parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// Add the template to the cache
	templateCache[t] = tmpl

	return nil
}
