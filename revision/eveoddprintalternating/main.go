package main

import (
	"context"
	"fmt"
	"sync"
)

const MAX_NUM = 100

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	el := make(chan struct{}, 1)
	ol := make(chan struct{}, 1)

	el <- struct{}{}

	es := generateEven(ctx, &wg, el, ol)
	ods := generateOdd(ctx, &wg, el, ol)

	processChannels(ctx, &wg, es, ods)

	wg.Wait()
}

func generateEven(ctx context.Context, wg *sync.WaitGroup, el chan struct{}, ol chan struct{}) <-chan int {
	stream := make(chan int)
	wg.Add(1)
	go func(){
		defer close(stream)
		defer wg.Done()
		for i := range(MAX_NUM/2 + 1){ // evens in [0, MAX_NUM]
			select{
			case <-el:
				stream <- i * 2
				ol <- struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}()

	return stream
}

func generateOdd(ctx context.Context, wg *sync.WaitGroup, el chan struct{}, ol chan struct{}) <-chan int {
	stream := make(chan int)
	wg.Add(1)
	go func(){
		defer close(stream)
		defer wg.Done()
		for i := range((MAX_NUM + 1) / 2){ // odds in [0, MAX_NUM]
			select{
			case <-ol:
				stream <- i * 2 + 1
				el <- struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}()

	return stream
}

func processChannels(ctx context.Context, wg *sync.WaitGroup, ec <-chan int, odc <-chan int) {
	wg.Add(1)
	go func(){
		defer wg.Done()
		chan1 := ec; chan2 := odc
		for chan1 != nil || chan2 != nil{
			select{
			case n1, ok := <-chan1:
				if(!ok) {chan1 = nil; continue}
				fmt.Println(n1)
			case n2, ok := <-chan2:
				if(!ok) {chan2 = nil; continue}
				fmt.Println(n2)
			case <-ctx.Done():
				return
			}
		}
	}()
}