package main

import (
	"fmt"
	"time"
)

//TODO: run the program and check that variable i
// was pinned for access from goroutine even after
// enclosing function returns.

func main() {
	for i := 1; i <= 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(5 * time.Second)
}
