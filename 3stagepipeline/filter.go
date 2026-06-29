package main

import (
	"context"
	"fmt"
	"sync"
)

func filter(ctx context.Context, input <-chan int, wg *sync.WaitGroup) <-chan int {
	output := make(chan int, 5)
	wg.Go(func(){
		defer close(output)
		for{
			select{
			case num, ok := <-input:
				if !ok {
					return 
				}
				if num % 2 == 0 {
					select{
					case output <- num:
					case <-ctx.Done():
						return 
					}
				}else{
					fmt.Println("skipping as its not a even")
				}
			case <-ctx.Done():
				return
			}
		}
	})
	return output
}