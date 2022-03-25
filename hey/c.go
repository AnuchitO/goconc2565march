package main

import "fmt"

func main() {
	c := make(chan string)

	go func() {
		msg := <-c
		fmt.Println("message: ", msg)
	}()

	c <- "gopher"
	fmt.Println("Done")
}
