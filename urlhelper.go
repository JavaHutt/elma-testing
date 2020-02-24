package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Analysed web address
type Analysed struct {
	url   string
	count int
}

// isValidUrl tests a string to determine if it is a well-structured URL or not.
func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

// SearchInURL counts the number of search word in a given URL.
func SearchInURL(url string) Analysed {
	if !isValidURL(url) {
		log.Fatal("Error! Not valid URL format!")
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return Analysed{
		url,
		strings.Count(string(body), searchWord),
	}
}
