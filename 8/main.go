package main

import (
	"fmt"
)

func main() {
	var number int64
	var bitPosition uint

	// Запрос ввода числа и позиции бита
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Print("Введите номер изменяемого бита: ")
	fmt.Scan(&bitPosition)

	// используем операции побитового ИЛИ
	// Установка i-го бита в 1
	number |= 1 << bitPosition
	fmt.Printf("Число в бинарном виде: %064b\n", uint64(number))
	fmt.Println("Число после установки бита в 1:", number)

	// Установка i-го бита в 0
	number &= ^(1 << bitPosition)
	fmt.Printf("Число в бинарном виде: %064b\n", uint64(number))
	fmt.Println("Число после установки бита в 0:", number)
}
