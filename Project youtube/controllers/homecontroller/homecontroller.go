package homecontroller

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles(
		"views/templates/mytemplate.html",
		"views/home/index.html",
	))
}

func Index(response http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(response, "layout", nil)
}
