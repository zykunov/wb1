package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mu sync.Mutex
var wg sync.WaitGroup

// создаем функцию накрутки счетчика
func counter(exitChan chan bool) {
	for {
		// будем останавливать рутину по передаче в канал
		select {
		default:
			//получаем эксклоюзивный доступ для переменной count
			mu.Lock()
			count++
			//разблокируем для других рутин
			mu.Unlock()
		case <-exitChan:
			// Получен сигнал об остановке
			fmt.Println("Остановлен")
			return
		}
	}
}

func main() {

	fmt.Println("Программа запущена")
	exitChan := make(chan bool)

	// создаем 5 рутин
	for i := 1; i <= 5; i++ {
		//добавляем в waitGroup
		wg.Add(1)
		//запускаем рутину
		go counter(exitChan)
		//указываем что рутина выполнилась
		wg.Done()
	}

	//Остановим программу через 3 секунды
	time.Sleep(time.Second * 3)
	//передаем в канал, чтоб остановить рутину
	exitChan <- true

	// ждем чтоб все рутины выполнились
	wg.Wait()
	//выводим
	fmt.Println("Итоговый счетчик:", count)
}
