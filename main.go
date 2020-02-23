package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	searchWord = "Go"
)

func searchInURL(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Count(string(body), searchWord)
}

func main() {
	fmt.Println(searchInURL("http://google.com"))
}
