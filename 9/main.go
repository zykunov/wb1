package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// создаем буфферизирорванный канал, с буфером такого же размера, как массив
	firstChannel := make(chan int, len(arr))
	// создаем второй такой же
	secondChannel := make(chan int, len(arr))

	// проходим по массиву
	for _, value := range arr {

		// пишем значение в канал
		firstChannel <- value
		// из канала во временную переменную
		tempVar := <-firstChannel

		// из временной переменной во второй канал и сразу умножаем
		secondChannel <- tempVar * 2
	}

	//еще раз проходимся и выводим значения из второго канала
	for _, _ = range arr {
		fmt.Println(<-secondChannel)
	}

}
