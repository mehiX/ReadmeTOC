package internal

import "net/url"

// ResponseData struct to pack the response
type ResponseData struct {
	URL   string `json:"url"`
	Toc   string `json:"toc"`
	Error string `json:"error"`
}

// Generator holds the url and the generated toc
type Generator struct {
	URL   url.URL
	ToC   string
	Error string

	Local bool
}
