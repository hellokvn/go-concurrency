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
	}

	for i, link := range links {
		sendRequest(i, link)
	}

	fmt.Println("Time:", time.Since(start)) // in average 5.203849865s
}

func sendRequest(i int, link string) {
	res, err := http.Get(link)

	if err != nil {
		fmt.Println(i, link, err)
		return
	}

	fmt.Println(i, link, res.StatusCode)
}
