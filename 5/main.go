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

	timer := time.NewTimer(time.Duration(seconds) * time.Second)

	ch := make(chan int)

	go func() {
		for {
			ch <- rand.Int()
		}
	}()

	for {
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
