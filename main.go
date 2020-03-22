package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mehiX/ReadmeTOC/toc"
)

var (
	help  = flag.Bool("help", false, "Print this message")
	serve = flag.String("serve", "", "Start a webserver on the specified address")
	path  = flag.String("path", "", "File path or URL to a markdown document")
)

/*
TODO
make it into a webserver if it receives a flag -listen with a port number
receive json or query param with url to README
respond with TOC

insert TOC under predefined tags or under "## Table of Contents" and return the full ReadME
*/

func init() {

	flag.Parse()

	if *help || ("" == *serve && "" == *path) {
		fmt.Fprintf(os.Stdout, "Usage: %s [-help] -serve ADDR -path URL\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

}

func main() {

	if "" != *serve {
		startServer()
	} else {
		generator := toc.NewGenerator(*path)

		generator.Generate()

		fmt.Fprintln(os.Stdout, generator.ToC)
	}
}

func startServer() {

	handler := mux.NewRouter()

	handler.HandleFunc("/", handleQueryParam).Methods(http.MethodGet)
	handler.HandleFunc("/json", handleJSON).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(*serve, handler))
}

func handleQueryParam(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("path")

	if "" == url {
		http.Error(w, "Missing query parameter: path", http.StatusBadRequest)
		return
	}

	generator := toc.NewGenerator(url)

	generator.Generate()

	w.Header().Set("Content-type", "text/plain")
	fmt.Fprint(w, generator.ToC)
}

type data struct {
	URL string `json:"url"`
	Toc string `json:"toc"`
}

func handleJSON(w http.ResponseWriter, r *http.Request) {

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
