package handlers

import (
	"html/template"
	"net/http"

	"github.com/mehiX/ReadmeTOC/internal"
)

// HandleQueryParam reads the query parameter (path) and
// responds with an html template
func HandleQueryParam(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("path")

	if "" == url {
		http.Error(w, "Missing query parameter: path", http.StatusBadRequest)
		return
	}

	generator := internal.NewGenerator(url)

	generator.Generate()

	w.Header().Set("Content-type", "text/html")

	d := internal.ResponseData{
		URL:   url,
		Toc:   generator.ToC,
		Error: generator.Error,
	}

	template.Must(template.ParseFiles("templates/home.html")).Execute(w, d)
}
