package main

import (
	"fmt"
	"net/http"
	"time"
)

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
		"https://yelp.com",
	}

	c := make(chan string)

	for i, link := range links {
		go sendRequest(i, link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	fmt.Println("Time:", time.Since(start)) // in average 1.111761763s
}

func sendRequest(i int, link string, c chan string) {
	res, err := http.Get(link)

	if err != nil {
		c <- fmt.Sprintf("%d | %s returns %d", i, link, -1)
		return
	}

	c <- fmt.Sprintf("%d | %s returns %d", i, link, res.StatusCode)
}
