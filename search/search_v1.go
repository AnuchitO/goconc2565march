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
		return Result(fmt.Sprintf("%s result for %q : duration %.2f \n", kind, query, d.Seconds()))
	}
}

// Google invokes Web, Image, and Video searches serially,
// appending them to the results slice.
func Google(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
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
