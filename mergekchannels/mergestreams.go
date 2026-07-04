package main

import (
	"context"
	"sync"
)

func MergeStreamsWithContext(ctx context.Context,streams ...<-chan int) chan int {
	junction := make(chan int)
	
	wg := &sync.WaitGroup{}
	for _, stream := range(streams){
		wg.Add(1)
		go func(st <-chan int){
			defer wg.Done()
			for{
				select{
				case x, ok := <-st:
					if !ok{
						return
					}
					select {
					case junction <- x:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return 
				}
			}
		}(stream)
	}

	go func(){
		defer close(junction)
		wg.Wait()
	}()

	return junction
}