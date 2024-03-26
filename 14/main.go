package main

import "fmt"

func main() {
	var a interface{} = 1
	var b interface{} = "1"
	var c interface{} = true
	var d interface{} = make(chan int)
	var e interface{} = 1.5

	getType(a)
	getType(b)
	getType(c)
	getType(d)
	getType(e)

}

func getType(variable interface{}) {
	// Определяем через switch
	switch variable.(type) {
	case int:
		fmt.Println("Это int")
	case string:
		fmt.Println("Это string")
	case bool:
		fmt.Println("Это bool")
	case chan int:
		fmt.Println("Это channel")
	case float64:
		fmt.Println("Это float")
	}
}
