package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

const (
	searchWord       = "Go"
	concurrencyLimit = 5
)

type totalCounter struct {
	increment func(n int)
	value     func() int32
}

// printResponse prints the counter for each URL and updates the total counter.
func printResponse(ch <-chan Analysed, total totalCounter) {
	for c := range ch {
		total.increment(c.count)
		fmt.Printf("Count for %s: %d\n", c.url, c.count)
	}
}

// countTotal is a simple counter with the ability to increment and to return value
func countTotal() totalCounter {
	total := new(int32)

	increment := func(n int) {
		atomic.AddInt32(total, int32(n))
	}
	value := func() int32 {
		return atomic.LoadInt32(total)
	}

	return totalCounter{
		increment,
		value,
	}
}

func main() {
	total := countTotal()
	wg := new(sync.WaitGroup)
	urlChan := make(chan Analysed)
	ctx := context.Background()
	sem := semaphore.NewWeighted(int64(concurrencyLimit))

	fmt.Println("Enter valid urls using space as delimeter")
	fmt.Println("Type 'quit' or tap Ctrl+C to stop and see the total counts")

	go printResponse(urlChan, total)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "quit" {
			break
		}
		sem.Acquire(ctx, 1)
		wg.Add(1)
		go func(url string, wg *sync.WaitGroup) {
			defer wg.Done()
			urlChan <- SearchInURL(url)
			sem.Release(1)
		}(text, wg)
	}
	wg.Wait()
	close(urlChan)
	fmt.Println("Total:", total.value())
}
