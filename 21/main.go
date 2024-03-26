package main

import "fmt"

/*
* группа классов, которые мы адаптируем
 */

// класс Английский язык
type English struct{}

// фраза на английском
func (english *English) sayHello() {
	fmt.Println("Hello world!")
}

// класс Русский язык
type Russian struct{}

// фраза на русском
func (russian *Russian) sayRussianHello(isCall bool) {
	if isCall {
		fmt.Println("Здарова мир!")
	}
}

/*
* адаптер
 */

// целевой интерфейс
type LanguageAdapter interface {
	Say()
}

// адаптер для английского языка
type EnglishAdapter struct {
	*English
}

// метод sayHello() английского языка
func (adapter *EnglishAdapter) Say() {
	adapter.sayHello()
}

// конструктор адаптера для англ языка
func NewEnglishAdapter(english *English) LanguageAdapter {
	return &EnglishAdapter{english}
}

// адаптер для русского языка
type RussianAdapter struct {
	*Russian
}

// метод sayRussianHello для русского языка
func (adapter *RussianAdapter) Say() {
	//тут адаптер по вызову
	adapter.sayRussianHello(true)
}

// конструктор адаптера для русского языка
func NewRussianAdapter(russian *Russian) LanguageAdapter {
	return &RussianAdapter{russian}
}

// Демонстрация
func main() {
	// массив с двумя элементами
	//первый элемент - адаптер английского языка.
	//второй - адаптер русского
	languages := [2]LanguageAdapter{NewEnglishAdapter(&English{}), NewRussianAdapter(&Russian{})}
	//
	for _, lang := range languages {
		// в каждом адаптере имеем возможность обратиться к речевым методам этих языков через общий Say()
		lang.Say()
	}
}
