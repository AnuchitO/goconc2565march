package main

import (
	"fmt"
	"math/rand"
	"time"
)

func heyMulti(msg string, c chan string) {
	defer close(c)
	for i := 0; i < 20; i++ {
		c <- fmt.Sprintf("hey %s Gopher %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go heyMulti("penguin", ch1)
	go heyMulti("gopher", ch2)

	for {
		select {
		case v := <-ch1:
			fmt.Printf("Wait ch1: %q\n", v)
		case v := <-ch2:
			fmt.Printf("Wait ch2: %q\n", v)
		}
	}

	fmt.Scanln()
	fmt.Println("Done")
}
