package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hey(msg string) {
	for i := 0; ; i++ {
		s := fmt.Sprintf("hey %s Gopher %d", msg, i)
		fmt.Println(s)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {

	go hey("world")
	

	go hey("penguin")

	fmt.Scanln()
	fmt.Println("Done")
}
