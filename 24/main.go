package main

import (
	"fmt"
	"math"
)

// свойства точки
type Point struct {
	x float64
	y float64
}

// конструтор точки
func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// нахождение расстояния
func GetLength(point1 *Point, point2 *Point) (result float64) {
	//вычисляем по формуле о длине вектора по координатам начала и конца
	result = math.Sqrt(math.Pow((point2.x-point1.x), 2) + math.Pow((point2.y-point1.y), 2))
	return result
}

func main() {

	point1 := NewPoint(1, 1)
	point2 := NewPoint(5, 5)

	fmt.Println(GetLength(point1, point2))
}
