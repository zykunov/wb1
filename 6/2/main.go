package main

import "fmt"

// закрытие канала через метод close
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	// на выходе возращаем канал(чтоб было что закрывать через close)
	return ch
}

func main() {
	number := Generator()
	fmt.Println(<-number)
	fmt.Println(<-number)
	close(number) // передаем переменную с каналом в close


