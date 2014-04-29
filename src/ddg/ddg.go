package main

import (
	"net/http"
	"net/url"
	"os"
	"regexp"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	isFeelingLucky := regexp.MustCompile("^(?:!(?:ducky)?\\s+|\\\\)(.*)")
	isDDG := regexp.MustCompile("^(!.*)")

	var outURL string

	switch {
	case isFeelingLucky.MatchString(query):
		match := isFeelingLucky.FindAllStringSubmatch(query, -1)[0][1]
		outURL = "https://encrypted.google.com/search?btnI=1&q=" + url.QueryEscape(match)
	case isDDG.MatchString(query):
		outURL = "https://duckduckgo.com/?q=" + url.QueryEscape(query)
	default:
		outURL = "https://encrypted.google.com/search?q=" + url.QueryEscape(query)
	}

	http.Redirect(w, r, outURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/search", handler)

	port := os.Args[1]
	http.ListenAndServe(":" + port, nil)
}
