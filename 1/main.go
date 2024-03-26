package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Height int
	Weight int
}

// Метод реализующий структуру human
func (h Human) GetWeight() int {
	return h.Weight
}

// Метод реализующий структуру human
func (h Human) GetHeight() int {
	return h.Height
}

type Action struct {
	Human //встраивание
	speed int
}

func (a Action) Run() string {
	name := a.Human.Name // обращаемся к свойству наследованному из структуры Human
	return "human:" + name + " is runing"
}

func main() {
	// создаем экземпляр структуры human
	firstHuman := Human{Age: 23, Name: "John Doe", Height: 183, Weight: 70}

	// вызываем методы структуры human
	fmt.Println("Height:", firstHuman.GetHeight())
	fmt.Println("Weight:", firstHuman.GetWeight())

	//cоздаем экземпляр структуры action
	secondHuman := Action{speed: 20, Human: firstHuman}

	//вызываем метод структуры action
	fmt.Println(secondHuman.Run())
}
