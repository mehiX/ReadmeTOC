package toc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var (
	titleMatcher     = regexp.MustCompilePOSIX("^[#]{1,} [A-Z0-9]{1,}")
	codeBlockMatcher = regexp.MustCompilePOSIX("^[`]{3}")
)

// NewGenerator creates a new generator based on a path
// the path can be local or a URL
func NewGenerator(path string) *Generator {

	pathURL, err := url.Parse(path)
	if nil != err {
		log.Panic(err)
	}

	return &Generator{
		URL:   *pathURL,
		Local: pathURL.Scheme == "",
	}
}

// Generate read the document and generate the ToC
func (g *Generator) Generate() {
	var r io.ReadCloser
	var err error

	var resp *http.Response

	if g.Local {
		r, err = os.Open(g.URL.String())
		if nil != err {
			g.Error = err.Error()
			return
		}
	} else {
		if resp, err = http.Get(g.URL.String()); nil == err {
			if http.StatusOK != resp.StatusCode {
				g.Error = fmt.Sprintf("Document not found: %s", g.URL.String())
				return
			}
			r = resp.Body
		}
	}

	defer r.Close()

	out := bytes.NewBuffer(*new([]byte))
	generateTOC(out, r)

	g.ToC = out.String()

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
