package main

import (
	"context"
	"sync"
)

const MAX_NUMBERS = 10000

func generator(ctx context.Context, wg *sync.WaitGroup) <-chan int {

	stream := make(chan int, 10)

	wg.Add(1)
	go func(ctx context.Context, stream chan int, wg *sync.WaitGroup){
		defer close(stream)
		defer wg.Done()

		for i := range(MAX_NUMBERS){
			select{
			case stream <- i:
			case <-ctx.Done():
				return
			}
		}

	}(ctx, stream, wg)

	return stream
}