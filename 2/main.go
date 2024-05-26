package main

import (
	"fmt"
	"sync"
)

func concurrentSquares(num int) {
	fmt.Println(num * num)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			concurrentSquares(num)
			wg.Done()
		}(num)
	}
	wg.Wait()
}
