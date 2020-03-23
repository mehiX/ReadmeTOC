package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mehiX/ReadmeTOC/toc"
)

type data struct {
	URL string `json:"url"`
	Toc string `json:"toc"`
}

func HandleJSON(w http.ResponseWriter, r *http.Request) {

	var d data

	if err := json.NewDecoder(r.Body).Decode(&d); nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	generator := toc.NewGenerator(d.URL)

	generator.Generate()

	d.Toc = generator.ToC

	w.Header().Set("Content-type", "application/json")

	if err := json.NewEncoder(w).Encode(d); nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
