package main

import "fmt"

func main() {
	c := make(chan string, 1)

	c <- "gopher"
	close(c)

	msg, ok := <-c
	fmt.Println("Ok:", ok)
	fmt.Println("message: ", msg)

	msg, ok = <-c
	fmt.Println("Ok:", ok)
	fmt.Println("message: ", msg)

	fmt.Println("Done")
}
