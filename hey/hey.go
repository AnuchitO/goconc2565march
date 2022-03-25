package main

import "fmt"

func hey(s string) {
	fmt.Println("hey Gopher", s)
}
func main() {
	go hey("A")

	fmt.Scanln()
	fmt.Println("Done")
}
