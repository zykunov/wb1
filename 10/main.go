package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// arr := []float64{-15.4, -32.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // другой набор, для теста

	tempGroups := make(map[string][]float64)

	for _, value := range arr {

		// смотрим - отрицательное или положительное значение?

		// округляем
		floor := int(math.Floor(value/10)) * 10

		// создаем итоговую мапу в подмножествами
		tempGroups[strconv.Itoa(floor)] = append(tempGroups[strconv.Itoa(floor)], value)

	}

	fmt.Println(tempGroups)
}
