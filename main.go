package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mehiX/ReadmeTOC/toc"
)

var (
	help = flag.Bool("help", false, "Print this message")
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

	if *help || len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stdout, "Usage: %s [-help] FILEPATH | URL\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

}

func main() {

	url := flag.Arg(0)
	generator := toc.NewGenerator(url)

	generator.Generate()

	fmt.Fprintln(os.Stdout, generator.ToC)

}
