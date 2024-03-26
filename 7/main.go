package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// создаем первую го рутину
func firstRoutine(m map[string]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	// включаем эксклюзивную блокировку т.к. будем править, общую для двух рутин, переменную
	mu.Lock()
	fmt.Println("Работает первая рутина")
	m["key"] = rand.Int()
	fmt.Println(m)
	//после правки выключаем переменную
	mu.Unlock()
	//отправляем в группу ожидания сигнал, что рутина отработала
	wg.Done()

}

// создаем вторую го рутину
func SecondRoutine(m map[string]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	fmt.Println("Работает вторая рутина")
	m["key"] = rand.Int()
	fmt.Println(m)
	mu.Unlock()
	wg.Done()

}

func main() {
	// mutex используем для создания эксклюзивного доступа к переменной
	var mu sync.Mutex

	//waitgroup используем т.к. программа выполняется очень быстро и го рутины не успевают запуститься
	var wg sync.WaitGroup

	myMap := make(map[string]int)

	//количество ожидаемых рутин известно - 2
	wg.Add(2)

	//запускаем рутины, передаем мап, указатель на мютекс и группу ожидания
	go firstRoutine(myMap, &mu, &wg)
	go SecondRoutine(myMap, &mu, &wg)

	wg.Wait()
}
