package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	var seconds int
	fmt.Print("через сколько секунд закончить?: ")
	fmt.Fscan(os.Stdin, &seconds)

	// создаем таймер, передаем время в него
	timer := time.NewTimer(time.Duration(seconds) * time.Second)

	ch := make(chan int)

	// запускаем рутину и постоянно передаем рандосмные числа
	go func() {
		for {
			ch <- rand.Int()
		}
	}()

	for {
		//смотрим что приходит в канал и если это от таймера, то прерываем программу
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("channel is closed!")
				return
			}
			fmt.Println(val)
		case <-timer.C:
			fmt.Println("time expired")
			return
		}

	}

}
