package main

import (
	"fmt"
	"math/rand"
	"time"
)

func heyc(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("hey %s Gopher %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go heyc("penguin", c)

	go func() {
		for i := 0; ; i++ {
			fmt.Printf("You say: %q\n", <-c)
		}
	}()

	fmt.Scanln()
	fmt.Println("Done")
}
