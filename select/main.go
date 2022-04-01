package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go count(1, 1, c1)
	go count(2, 2, c2)

	for {
		select {
		case s1 := <-c1:
			fmt.Println(s1)
		case s2 := <-c2:
			fmt.Println(s2)
		}
	}

}

func count(id int, seconds int, c chan<- string) {
	for {
		time.Sleep(time.Second * time.Duration(seconds))
		c <- fmt.Sprintf("receive from channel %d", id)
	}
}
