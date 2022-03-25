package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		rand.Seed(time.Now().UnixNano())
		d := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(d)
		return Result(fmt.Sprintf("%s result for %q : duration %d \n", kind, query, d.Milliseconds()))
	}
}

// Run the Web, Image, and Video searches concurrently, and wait for all results.
// No locks. No condition variables. No callbacks.
func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
