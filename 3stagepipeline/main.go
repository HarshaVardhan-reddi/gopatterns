package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("3 stage pipeline where generate -> modify -> filter will be happening")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// generates the stream
	genstream := generator(ctx, wg)

	// modifies the stream
	modified := modifier(ctx, genstream, wg)

	// filter stream
	filtered := filter(ctx, modified, wg)

	// print filtered
	printFiltered(ctx, filtered, wg)

	wg.Wait()
}

func printFiltered(ctx context.Context, stream <-chan int, wg *sync.WaitGroup) {
	wg.Go(func() {
		for {
			select {
			case ele, ok := <-stream:
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
