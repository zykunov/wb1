package main

import "fmt"

func binarySearch(needle int, haystack []int) bool {

	low := 0
	//количество элементов
	high := len(haystack) - 1

	// проходим по массиву,, по возрастанию
	for low <= high {

		//находим каждый раз середину массива
		median := (low + high) / 2

		//если средний элемент массива меньше искомого
		if haystack[median] < needle {
			// то ставим начало на середину массива, чтоб в следущий раз начинать проход с середины
			low = median + 1
		} else {
			// если больше то наоборот верхнюю границу массива сдвигаем на середину
			high = median - 1
		}
	}

	// проверка, если нижняя груница равна длине массива или нижний элемент не равен искомому -
	// то искомого элемента в массиве нет
	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(70, items))
}
