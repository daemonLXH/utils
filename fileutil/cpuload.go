package main

import (
	"flag"
	"sync"
)

func add(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += i
	}
	return n
}

func main() {
	var count int
	var wg sync.WaitGroup

	flag.IntVar(&count, "c", 1000, "-c")

	flag.Parse()

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				add(i)
			}

		}()

	}

	wg.Wait()
}
