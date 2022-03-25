package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(6)
	for i := 0; i <= 5; i++ {
		go func(x int) {
			defer wg.Done()
			fmt.Println("long procesing..")
			time.Sleep(1 * time.Second)
			fmt.Println(x)
		}(i)
	}

	wg.Wait()
	fmt.Println("Done")
}
