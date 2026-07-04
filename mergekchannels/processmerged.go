package main

import (
	"context"
	"fmt"
	"sync"
)

func ProcessMergedStreamWithContext(ctx context.Context,stream <-chan int){
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func (){
		defer wg.Done()
		for{
			select{
			case x,ok := <-stream:
				if !ok{return}
				fmt.Println(x)
			case <-ctx.Done():
				return
			}
		}
	}()
	wg.Wait()
}