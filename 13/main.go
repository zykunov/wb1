package main

import "fmt"

func main() {

	// стартуем переменные
	first := 38
	second := 45

	fmt.Println(first, second)

	// первая переменная равна сумме переменных
	first = first + second //83
	// вторая равна 83 - 45 = 38 - уже поменялась с первой
	second = first - second //38
	// и первая 83 - 38 = 45 - поменялась со второй переменной
	first = first - second
	fmt.Println(first, second)

}
