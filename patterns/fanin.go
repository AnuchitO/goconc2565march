package main

import (
	"fmt"
	"time"
)

func say(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(1e3) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			v := <-input1
			c <- v
		}
	}()
	go func() {
		for {
			v := <-input2
			c <- v
		}
	}()
	return c
}

func main() {
	gch := say("gopher")
	pch := say("penguin")

	c := fanIn(gch, pch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("gopher what's up.")
}
