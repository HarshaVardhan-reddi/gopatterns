package main

import (
	"context"
	"sync"
)

func modifier(ctx context.Context,inpstream <-chan int, wg *sync.WaitGroup) <-chan int {
	op := make(chan int, 5)
	wg.Add(1)

	go func(inp <-chan int, output chan<- int, wg *sync.WaitGroup){
		defer wg.Done()
		defer close(output)

		for{
			select{
			case num, ok := <-inp:
				if(!ok){
					return
				}
				select{
				case output <-(num * num):
				case <-ctx.Done():
					return 
				}
			case <-ctx.Done():
				return
			}
		}

	}(inpstream, op, wg)

	return op
}