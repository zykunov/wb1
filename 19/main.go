package main

import (
	"fmt"
	"strings"
)

func main() {

	// слово будем хранить через strings.Builder
	var sb strings.Builder

	//так как у нас unicode, то используем руны
	word := []rune("главрыба")
	fmt.Println("Исходное слово:", string(word))

	//замеряем длину
	runesNumber := len(word)

	//проходим от последнего символа к первому
	for i := runesNumber - 1; i >= 0; i-- {
		//записываем
		sb.WriteRune(word[i])
	}

	//выводим
	fmt.Println("Перевернутое:", sb.String())
}
