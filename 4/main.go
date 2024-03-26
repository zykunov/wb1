package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

/* запуск программы - go run .\4\main.go 4, где 4 это количество воркеров, цифра может быть любой */

// делаем waitgroup для разруливания конкурентности
var wg sync.WaitGroup

// создаем worker
func worker(number int, ch <-chan int) {
	for {
		// выводим сообщение из канала
		fmt.Printf("data from channel %v, worker number %v\n", ch, number)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Укажите количество воркеров - запуск программы - go run main.go 4, где 4 это количество воркеров")
		return
	}

	// определяем количество воркеров
	workersNumber, _ := strconv.Atoi(os.Args[1])
	fmt.Println("count of workers:", workersNumber)

	// основной канал, куда будем передавать данные
	ch := make(chan int)
	// var ch chan int

	// 1 cоздаем воркеры для чтения и вывода
	for i := 1; i <= workersNumber; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, ch)
		}(i)
	}

	// 2 создаем канал на системный вызов, будем использовать для ожидания ctrl+c
	exitCh := make(chan os.Signal, 1)
	// попдисываемся на событие завершения
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	// рутина для закрытия основного канала
	go func() {
		// ждем сообщение в канал закрытия
		<-exitCh
		// закрываем канал
		close(ch)
	}()

	// 3 запускаем рутину для записи в основной канал
	go func() {
		for {
			ch <- rand.Int()
			time.Sleep(time.Millisecond * 2)
		}
	}()

	wg.Wait()
}
