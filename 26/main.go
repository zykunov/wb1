package main

import (
	"fmt"
	"strings"
)

func main() {
	// строка с повтором
	stringsVar := "abcdeAB"

	//строка без повтора
	// stringsVar := "abcdefghijklmn"

	fmt.Println("строка:", stringsVar)

	// мапа, чтоб проверять есть ли в ней такой ключ
	charMap := make(map[string]int)
	//флаг
	var flag bool

	//проходим по строке
	for i := 0; i <= len(stringsVar)-1; i++ {

		//чтоб не зависить от регистра - приводим все к нижнему
		char := strings.ToLower(string(stringsVar[i]))
		//проверяем нет ли уже такого символа в мапе
		_, ok := charMap[char]
		if ok {
			//если есть то ставим флаг
			flag = true
		}
		charMap[char] = 1

	}

	// если flag == true
	if flag {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}

}
