package main

import (
	"context"
	"sync"
)

func GenerateEvenNum(ctx context.Context, wg *sync.WaitGroup, signal1 chan struct{}, signal2 chan struct{}) <-chan int {
	stream := make(chan int)
	wg.Add(1)
	go func(){
		defer close(stream)
		defer wg.Done()
		for i := range(MAXNUM / 2) {
			select{
			case <-signal1:
				stream <- i * 2
				signal2 <- struct{}{}
			case <-ctx.Done():
				return 
			}
		}
	}()
	return stream
}