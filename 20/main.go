package main

import (
	"fmt"
	"strings"
)

func main() {
	words := "snow dog sun"
	fmt.Println("Изначальная строка:", words)

	// с помощью strings.Fields строка преобразуется в массив, разделитель - пробелы
	wordsArr := strings.Fields(words)

	fmt.Printf("%s ", "Перевернутая строка:")

	// проходим по массиву в обратном порядке и выводим слова
	for i := len(wordsArr) - 1; i >= 0; i-- {
		fmt.Printf("%s ", wordsArr[i])
	}

}
