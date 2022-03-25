package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sendAll(msg string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("hey %s Gopher %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
	close(c)
}

func main() {
	c := make(chan string, 10)

	go sendAll("penguin", c)

	go func() {
		for v := range c {
			fmt.Print("Wait : ")
			fmt.Printf("You say: %q\n", v)
		}
	}()

	fmt.Scanln()
	fmt.Println("Done")
}
