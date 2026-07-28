package main

import(
	"context"
	"sync"
)


func GenerateOddNum(ctx context.Context, wg *sync.WaitGroup, signal1 chan struct{}, signal2 chan struct{}) <-chan int {
	stream := make(chan int)
	wg.Add(1)
	go func(){
		defer close(stream)
		defer wg.Done()
		for i := range(MAXNUM / 2) {
			select{
			case <-signal2:
				stream <- i * 2 + 1
				signal1 <- struct{}{}
			case <-ctx.Done():
				return 
			}
		}
	}()
	return stream
}