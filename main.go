package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	titleMatcher     *regexp.Regexp
	codeBlockMatcher *regexp.Regexp
	url              string

	help = flag.Bool("help", false, "Print this message")
)

func init() {

	flag.Parse()

	if *help || len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stdout, "Usage: %s [-help] FILEPATH | URL\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	var err error

	url = flag.Arg(0)

	titleMatcher, err = regexp.CompilePOSIX("^[#]{1,} [A-Z0-9]{1,}")
	if nil != err {
		log.Panicln(err)
	}

	codeBlockMatcher, err = regexp.CompilePOSIX("^[`]{3}")
	if nil != err {
		log.Panicln(err)
	}
}

func main() {

	var r io.ReadCloser
	var err error

	var resp *http.Response

	if resp, err = http.Get(url); nil == err {
		if http.StatusOK != resp.StatusCode {
			log.Panicf("Document not found: %s", url)
		}
		r = resp.Body
	} else {
		fmt.Fprintf(os.Stderr, "URL error: %v", err)
		// open the file
		r, err = os.Open(url)
		if nil != err {
			log.Panicln(err)
		}
	}

	defer r.Close()

	generateTOC(os.Stdout, r)
}

func generateTOC(w io.Writer, r io.Reader) {
	scanner := bufio.NewScanner(r)

	// holds the current count for each indentation level
	counter := make([]int, 6)

	// will be true each time we are in a code block
	// this is to skip comment lines (also starting with #) in code blocks
	codeBlock := false
	for scanner.Scan() {
		line := scanner.Text()
		codeBlock = codeBlock != codeBlockMatcher.MatchString(line)
		if !codeBlock && titleMatcher.MatchString(line) {
			anchor := getAnchor(line)

			parts := strings.SplitN(line, " ", 2)
			counter[len(parts[0])]++

			fmt.Fprintf(w, "%s [%s](#%s)\n", fmt.Sprintf("%s%d.", getIndent(len(parts[0])), counter[len(parts[0])]), parts[1], anchor)

			// clear counters
			for i := len(parts[0]) + 1; i < len(counter); i++ {
				counter[i] = 0
			}
		}
	}

}

func getIndent(pos int) string {
	str := ""
	for i := 1; i < pos; i++ {
		str += "    "
	}

	return str
}

func getAnchor(txt string) string {
	chars, _ := regexp.CompilePOSIX("[^a-zA-Z0-9 ]*")

	res := chars.ReplaceAll([]byte(txt), []byte(""))

	str := strings.Trim(string(res), " ")
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "-")

	return str
}
