package main

import (
	"context"
	"fmt"
	"time"
)

// Завершение рутины с помощью контекста
func main() {
	// создаем канал
	forever := make(chan struct{})

	// создаем контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())

	// передаем контекст в рутину
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // if cancel() execute
				forever <- struct{}{}
				return
			default:
				fmt.Println("for loop")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	<-forever
	fmt.Println("finish")
}
