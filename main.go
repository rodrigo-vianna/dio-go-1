package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
	}
}

func main() {
	c := make(chan string)

	go ping(c)
	go pong(c)

	go func() {
		for {
			fmt.Println(<-c)
			time.Sleep(1 * time.Second)
		}
	}()
	fmt.Scanln()
}
