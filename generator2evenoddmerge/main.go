package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup

	even := generateEven(ctx, &wg, 1000)
	odd := generateOdd(ctx, &wg, 1000)

	channels := []<-chan int{even, odd}

	merged := mergeAndQualifyNums(ctx, &wg, channels)

	process(ctx, &wg, merged)

	wg.Wait()
}

func generateEven(ctx context.Context, wg *sync.WaitGroup, maxnum int) <-chan int {
	es := make(chan int, 1)

	wg.Go(func() {
		defer close(es)
		for i := range maxnum / 2 {
			select {
			case es <- i * 2:
			case <-ctx.Done():
				return
			}
		}
	})

	return es
}

func generateOdd(ctx context.Context, wg *sync.WaitGroup, maxnum int) <-chan int {
	ods := make(chan int, 1)

	wg.Go(func() {
		defer close(ods)
		for i := range maxnum / 2 {
			select {
			case ods <- i*2 + 1:
			case <-ctx.Done():
				return
			}
		}
	})

	return ods
}

func mergeAndQualifyNums(ctx context.Context, wg *sync.WaitGroup, streams []<-chan int) <-chan int {
	filtered := make(chan int, 1)

	wg.Go(func() {
		innerwg := &sync.WaitGroup{}
		defer close(filtered)
		for _, stream := range streams {
			innerwg.Add(1)
			
			go func(st <-chan int) {
				defer innerwg.Done()
				for {
					select {
					case num, ok := <-st:
						if(!ok) {return}
						if num > 10 {
							filtered <- num
						}
					case <-ctx.Done():
						return
					}
				}
			}(stream)
		}

		innerwg.Wait()
	})

	return filtered
}

func process(ctx context.Context, wg *sync.WaitGroup, result <-chan int) {
	wg.Go(func() {
		for {
			select {
			case ele, ok := <-result:
				if !ok {
					return
				}
				fmt.Println(ele)
			case <-ctx.Done():
				return
			}
		}
	})
}
