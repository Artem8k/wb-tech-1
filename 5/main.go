package main

import (
	"fmt"
	"log"
	"time"
)

func chanReader(d chan string, s chan bool) {
	for {
		select {
		case data, ok := <-d:
			fmt.Printf("Получил данные: %s", data)

			if !ok {
				return
			}

		case <-s:
			return
		}
	}

}

func main() {

	var sec int
	fmt.Println("Введите количество секунд: ")
	_, err := fmt.Scanln(&sec)

	if err != nil {
		log.Fatal(err)
	}

	d := make(chan string)
	s := make(chan bool)

	go chanReader(d, s)

	go func() {
		for {
			var data string
			fmt.Println("Введите данные: ")
			fmt.Scanln(&data)
			d <- data
		}
	}()

	time.Sleep(time.Duration(sec) * time.Second)
}
