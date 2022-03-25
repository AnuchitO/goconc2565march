package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	go func(a, b int) {
		c := a + b
		ch <- c
		close(ch)
	}(1, 2)

	r, ok := <-ch
	fmt.Println("Ok:", ok)
	fmt.Printf("computed value %v\n", r)

	// ch <- 55
	// r, ok = <-ch
	// fmt.Println("Ok:", ok)
	// fmt.Printf("computed value %v\n", r)
}
