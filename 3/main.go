package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	arr := [5]int{2, 4, 6, 8, 10}

	var TotalRes int
	waitGroup.Add(5) // т.к. известно количество итераций, заранее объявляем количество ожидаемых.
	for _, value := range arr {
		go func(i int) {
			result := i * i    // находим квадрат числа
			TotalRes += result // прибавляем в общую суммц
			waitGroup.Done()
		}(value)
	}
	waitGroup.Wait()
	fmt.Println(TotalRes)
}
