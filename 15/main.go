package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}*/

/* 1 Негативные последствия - в переменной v хранится большая строка, далее мы помещаем в переменную justString срез на 100 элементов
из этой строки. При этом, поскольку срез это лишь указатель на ячейку в памяти, то переменная v не сможет быть очищена сборщиком мусора
так как на нее как раз и указывает слайс.
*/

/*2 решение
Во первых, для создания больших строк в createHugeString можно использовать strings.Builder
Во вторых, скопировать строку нужной длины в отдельную переменную, с отдельным выделением памяти под это*/

var justString string

func main() {

	someFunc(&justString)
	fmt.Println(justString)
}

func someFunc(justString *string) {
	v := createHugeString(1 << 10)
	// Копируем строчку через strings.Clone, тем самым веделится новый(меньший) участок памяти
	*justString = strings.Clone(v)[:100]
}

func createHugeString(size int) string {
	// создаем новый генератор
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	//для создания больших строк используем  strings.Builder
	hugeString := strings.Builder{}
	for i := 0; i < size; i++ {
		hugeString.WriteRune('a' + rune(generator.Intn('z'-'a'+1)))
	}
	return hugeString.String()
}
