package main

import (
	"fmt"
	"math/big"
)

func main() {
	// создаем большие числа с помощью big.NewInt
	a := big.NewInt(5000000000000)
	b := big.NewInt(5000000000000)

	//вызов функций
	fmt.Println(add(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(mult(a, b))
	fmt.Println(div(a, b))

}

// сложение
func add(a *big.Int, b *big.Int) (res *big.Int) {
	// создаем новый элемент big.Int в переменной res
	res = new(big.Int)
	// с помощью метода Add библиотеки bit.Int выполняем операцию
	return res.Add(a, b)
}

// вычитание
func sub(a *big.Int, b *big.Int) (res *big.Int) {
	res = new(big.Int)
	return res.Sub(a, b)
}

// умножение
func mult(a *big.Int, b *big.Int) (res *big.Int) {
	res = new(big.Int)
	return res.Mul(a, b)
}

// деление
func div(a *big.Int, b *big.Int) (res *big.Int) {
	res = new(big.Int)
	return res.Div(a, b)
}
