package main

import (
	"fmt"
	"sync"
)

func main() {

	var waitGroup sync.WaitGroup

	arr := [5]int{2, 4, 6, 8, 10}

	waitGroup.Add(5) // пришлось использовать wait группы, т.к. слишком быстро выполнялась программа и ничего не выводилось
	for _, value := range arr {

		go func(i int) {
			result := i * i
			fmt.Println(result)
			waitGroup.Done()
		}(value)
	}
	waitGroup.Wait()
}
