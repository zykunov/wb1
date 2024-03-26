package main

import (
	"fmt"
	"time"
)

func mySleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	fmt.Println("Первый вывод")
	fmt.Println("Засыпаем на 2 секунды")
	mySleep(time.Duration(2 * time.Second))
	fmt.Println("Второй вывод")
}
