package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)

		// TODO: send all values on channel without blocking
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	// TODO: print all of the values received on the channel
}
