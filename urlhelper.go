package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Analysed web address
type Analysed struct {
	url   string
	count int
	err   error
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
func SearchInURL(url string) (result Analysed) {
	result.url = url

	if !isValidURL(url) {
		result.err = errors.New("Error! Not valid URL format")
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		result.err = err
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.err = err
		return
	}

	result.count = strings.Count(string(body), searchWord)

	return
}
