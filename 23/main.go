package main

import "fmt"

func main() {

	arr := []string{"A", "B", "C", "D", "E"}
	i := 2

	fastDelete(arr, i)

	slowDelete(arr, i)

}

func fastDelete(arr []string, i int) {

	// копируем последний элемент на место удаляемого
	arr[i] = arr[len(arr)-1]

	// записываем пустоту в песледний элемент
	arr[len(arr)-1] = ""

	// Подрезаем слайс
	arr = arr[:len(arr)-1]

	fmt.Println(arr)
}
