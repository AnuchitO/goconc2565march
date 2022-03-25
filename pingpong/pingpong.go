package main

import (
	"fmt"
	"time"
)

type Ball struct{ hits int }

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, "hit", ball.hits, "times")
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

func main() {
	table := make(chan *Ball)
	go player("gopher", table)
	go player("penguin", table)

	table <- &Ball{}
	time.Sleep(2 * time.Second)
	ball := <-table

	fmt.Println("Total hit: ", ball.hits)
}
