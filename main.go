package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mehiX/ReadmeTOC/toc"
	"github.com/mehiX/ReadmeTOC/toc/handlers"
)

var (
	help  = flag.Bool("help", false, "Print this message")
	serve = flag.String("serve", "", "Start a webserver on the specified address")
	path  = flag.String("path", "", "File path or URL to a markdown document")
)

/*
TODO
add homepage

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

	handler.HandleFunc("/", handlers.Home).Methods(http.MethodGet)
	handler.HandleFunc("/query", handlers.HandleQueryParam).Methods(http.MethodGet)
	handler.HandleFunc("/json", handlers.HandleJSON).Methods(http.MethodGet)

	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(*serve, handler))
}
