package render

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var parsedTemplate, _ = template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		log.Println("Error parsing template: ", err)
	}
}
