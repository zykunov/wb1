package main

import "fmt"

// делаем функцию для проверки существования элемента в массиве
func contains(arr []int, findValue int) {
	for _, value := range arr {
		if value == findValue {
			fmt.Println(value)
		}
	}
}

func main() {

	// хардкодим множества
	firstArr := []int{2, 56, 34, 84, 25, 21, 245, 673, 62}
	secondArr := []int{2, 6, 56, 245, 72, 92, 31}

	// в цикле проходим по массиву и каждый элемент отправляем в самописную функцию
	for _, value := range firstArr {
		//возвращает все пересечения- 2, 56, 245
		contains(secondArr, value)
	}

}
