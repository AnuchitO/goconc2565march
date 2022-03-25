package main

import "fmt"

func main() {
	owner := func() <-chan int {
		ch := make(chan int)
		go func(c chan<- int) {
			defer close(c)
			for i := 0; i < 5; i++ {
				fmt.Printf("Sending: %d\n", i)
				ch <- i
			}
		}(ch)
		return ch
	}

	consumer := func(ch <-chan int) {
		// read values from channel
		for v := range ch {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("Done receiving!")
	}

	ch := owner()
	consumer(ch)
}
