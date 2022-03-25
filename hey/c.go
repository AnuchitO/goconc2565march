package main

import "fmt"

func main() {
	c := make(chan string)

	go func() {
		c <- "gopher"
	}()

	msg := <-c
	fmt.Println("message: ", msg)
	fmt.Println("Done")
}
