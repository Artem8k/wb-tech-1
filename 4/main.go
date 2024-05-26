package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(id int, data chan string, wg *sync.WaitGroup, stop chan os.Signal) {
	defer wg.Done()
	for {
		select {
		case d, ok := <-data:
			if !ok {
				return
			}
			fmt.Printf("Воркер %d получил: %s\n", id, d)
		case <-stop:
			return
		}
	}
}

func main() {
	var workers int

	fmt.Print("Введите количество воркеров: ")
	_, err := fmt.Scanln(&workers)

	if err != nil {
		log.Fatal(err)
	}

	d := make(chan string)
	stop := make(chan os.Signal)
	var wg sync.WaitGroup

	// Создаем и запускаем воркеры
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, d, &wg, stop)
	}

	// Главная горутина записывает данные в канал
	go func() {
		for {
			var data string
			fmt.Println("Введите данные: ")
			fmt.Scanln(&data)
			d <- data
		}
	}()

	// Обработка сигнала завершения по Ctrl+C
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalCh
		close(stop)
	}()

	// Ожидание завершения работы воркеров
	wg.Wait()
}
