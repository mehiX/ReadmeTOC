package handlers

import (
	"html/template"
	"net/http"

	"github.com/mehiX/ReadmeTOC/toc"
)

func HandleQueryParam(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("path")

	if "" == url {
		http.Error(w, "Missing query parameter: path", http.StatusBadRequest)
		return
	}

	generator := toc.NewGenerator(url)

	generator.Generate()

	w.Header().Set("Content-type", "text/html")

	d := Data{
		URL: url,
		Toc: generator.ToC,
	}

	template.Must(template.ParseFiles("tmpl/home.html")).Execute(w, d)
}
