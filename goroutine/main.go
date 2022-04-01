package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()

	links := []string{
		"https://go.dev",
		"https://medium.com",
		"https://wikipedia.com",
		"https://google.com",
		"https://instagram.com",
		"https://duckduckgo.com",
		"https://youtube.com",
		"https://heroku.com",
		"https://amazon.com",
		"https://facebook.com",
		"https://twitter.com",
	}

	for i, link := range links {
		wg.Add(1) // incease WaitGroup counter
		go sendRequest(i, link)
	}

	wg.Wait() // waits until WaitGroup is <= 0

	fmt.Println("Time:", time.Since(start)) // in average 840.856718ms
}

func sendRequest(i int, link string) {
	defer wg.Done() // incease WaitGroup counter

	res, err := http.Get(link)

	if err != nil {
		fmt.Println(i, link, err)
		return
	}

	fmt.Println(i, link, res.StatusCode)
}
