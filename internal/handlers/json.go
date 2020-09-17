package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mehiX/ReadmeTOC/internal"
)

// HandleJSON receives a JSON object in the form {"url": ""} and returns {"url": "...", "toc": "..."}
func HandleJSON(w http.ResponseWriter, r *http.Request) {

	var d internal.ResponseData

	if err := json.NewDecoder(r.Body).Decode(&d); nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	generator := internal.NewGenerator(d.URL)

	generator.Generate()

	d.Toc = generator.ToC
	d.Error = generator.Error

	w.Header().Set("Content-type", "application/json")

	if err := json.NewEncoder(w).Encode(d); nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
