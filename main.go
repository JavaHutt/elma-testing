package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	searchWord = "Go"
)

// isValidUrl tests a string to determine if it is a well-structured url or not.
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

//searchInURL counts the number of search word
//in a given URL
func searchInURL(url string) int {
	if !isValidURL(url) {
		log.Fatal("Error! Not valid url format!")
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

	return strings.Count(string(body), searchWord)
}

func printResponse(ch chan int, total *int) {
	for c := range ch {
		*total += c
		fmt.Printf("Count for test %d\n", c)
	}
}

func main() {
	var total int
	fmt.Println("Enter valid urls using space as delimeter")

	urlChan := make(chan int)
	go printResponse(urlChan, &total)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		text := scanner.Text()
		urlChan <- searchInURL(text)
	}
	fmt.Println("Total:", total)
}
