package main

import (
	"fmt"
	"sync"
)

func concurrentSquares(num int) int {
	return num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	var summ int

	for _, num := range numbers {
		wg.Add(1)

		go func(n int) {
			summ += concurrentSquares(n)
			wg.Done()
		}(num)

	}
	wg.Wait()

	fmt.Println(summ)
}
