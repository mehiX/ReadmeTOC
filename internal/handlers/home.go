package handlers

import (
	"html/template"
	"net/http"
)

// Home handler that returns the homepage based on an html template
func Home(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("templates/home.html")).Execute(w, nil)
}
