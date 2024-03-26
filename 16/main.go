package main

import (
	"fmt"
	"math/rand"
)

// функция генерации массива со случайными цифрами
func generateArr() (arr []int) {

	// создаем массив на 10к элементов, чтоб было что посортировать
	arr = make([]int, 10000)

	//заполняем массив
	for i := 0; i <= 9999; i++ {
		arr[i] = rand.Intn(10000)
	}

	return arr
}

// функция сортировки
func qsort(a []int) []int {
	//проверяем длину массива, т.к. он будет рекурсивно дробиться на 2 части
	if len(a) < 2 {
		return a
	}

	//
	left, right := 0, len(a)-1
	// выбираем опорный элемент, рандомно в пределах размера массива
	pivotIndex := rand.Int() % len(a)
	fmt.Println(pivotIndex)

	//a[pivotIndex] = последний элемент массива, a[right] = rand.Int() % len(a)
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Проходим по каждому элементу
	for i := range a {
		//если текущий элемент меньше опорного элемента,то
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	// перемещаем опорный элемента, после последнего меньшего элемента
	a[left], a[right] = a[right], a[left]
	// рекурсивно вызываем эту функцию для каждой части и так далее
	qsort(a[:left])
	qsort(a[left+1:])
	return a
}

func main() {

	// проверяем на маленьком массиве
	// fmt.Println(qsort([]int{5, 6, 6, 7, 98, 1, 315, 3, 5, 7}))

	// проверка на большом массиве
	fmt.Println(qsort(generateArr()))
}
