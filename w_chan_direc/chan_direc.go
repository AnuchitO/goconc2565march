package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func sendMessage(ch1 chan<- string) {
	for i := 0; i < 5; i++ {
		ch1 <- fmt.Sprintf("Sending: %d\n", i)
	}
}

func relayMessage(ch1 <-chan string, ch2 chan<- string) {
	for v := range ch1 {
		ch2 <- v
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendMessage(ch1)
	go relayMessage(ch1, ch2)

	for v := range ch2 {
		fmt.Printf("Received: %s\n", v)
	}
}
