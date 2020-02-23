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
	"sync"
)

const (
	searchWord = "Go"
)

type analysed struct {
	url   string
	count int
}
type totalCounter struct {
	increment func(n int)
	value     func() int
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

// searchInURL counts the number of search word in a given URL.
func searchInURL(url string) analysed {
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

	return analysed{
		url,
		strings.Count(string(body), searchWord),
	}
}

// printResponse prints the counter for each URL and updates the total counter.
func printResponse(ch <-chan analysed, total totalCounter, mutex *sync.Mutex) {
	for c := range ch {
		mutex.Lock()
		total.increment(c.count)
		mutex.Unlock()
		fmt.Printf("Count for %s: %d\n", c.url, c.count)
	}
}

// countTotal is a simple counter with the ability to increment and to return value
func countTotal() totalCounter {
	total := 0

	return totalCounter{
		func(n int) {
			total += n
		},
		func() int {
			return total
		},
	}
}

func main() {
	total := countTotal()
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	urlChan := make(chan analysed)

	fmt.Println("Enter valid urls using space as delimeter")
	fmt.Println("Type 'quit' or tap Ctrl+C to stop and see the total counts")

	go printResponse(urlChan, total, m)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "quit" {
			break
		}
		wg.Add(1)
		go func(url string, wg *sync.WaitGroup) {
			urlChan <- searchInURL(url)
			wg.Done()
		}(text, wg)
	}
	wg.Wait()
	close(urlChan)
	fmt.Println("Total:", total.value())
}
