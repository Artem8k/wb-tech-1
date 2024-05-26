package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	i int
	m sync.Mutex
}

func (c *Counter) Increment() {
	c.m.Lock()
	defer c.m.Unlock()
	c.i++
	fmt.Println(c.i)
}

func worker(c *Counter) {
	for i := 0; i < 10; i++ {
		c.Increment()
	}
}

func main() {
	var wg sync.WaitGroup
	c := Counter{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			worker(&c)
			wg.Done()
		}()
	}
	wg.Wait()
}
