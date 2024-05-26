package main

import (
	"fmt"
	"sync"
)

func squareNum(num int, squaredX chan int) {
	squaredX <- num * num
}

func read(squaredX chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case num, ok := <-squaredX:
			if !ok {
				return
			}
			fmt.Println(num)
		default:
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup

	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var x = make(chan int, 10)
	var squaredX = make(chan int, 10)

	// записываю в канал числа из массива
	for _, num := range nums {
		x <- num
	}
	close(x)

	// для вычисления каждого квадрата читаю из канала x и создаю новую горутину для каждого вычисления
	for num := range x {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			squareNum(num, squaredX)
		}(num)
	}

	// в отдельной горутине запускаю чтение из канала squaredX и вывод в консоль
	wg.Add(1)
	go read(squaredX, &wg)

	wg.Wait()
	close(squaredX)
}
