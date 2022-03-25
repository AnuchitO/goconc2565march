package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(c chan<- int) {
		defer close(c)
		for i := 0; i < 5; i++ {
			fmt.Printf("Sending: %d\n", i)
			c <- i
		}
	}(ch)

	go func(c <-chan int) {
		for v := range c {
			fmt.Printf("Received: %d\n", v)
		}
	}(ch)

	time.Sleep(5 * time.Second)
}
