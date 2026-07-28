package main

import (
	"context"
	"fmt"
	"sync"
)

const MAXNUM = 10

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sig1 := make(chan struct{}, 1)
	sig2 := make(chan struct{}, 1)
	defer cancel()
	wg := &sync.WaitGroup{}

	evenstream := GenerateEvenNum(ctx, wg, sig1, sig2)
	oddstream := GenerateOddNum(ctx, wg, sig1, sig2)

	streams := make([]<-chan int, 0)
	streams = append(streams, evenstream)
	streams = append(streams, oddstream)

	sig1 <- struct{}{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		closedstreams := 0
		var evenclosed bool = false
		var oddclosed bool = false
		allclosed := func() bool {
			if closedstreams == 2 {
				return true
			} else {
				return false
			}
		}
		for {
			if !evenclosed {
				ele, ok := <-evenstream
				if(ok){fmt.Println(ele)}
				if !ok {
					closedstreams += 1
					evenclosed = true
					if allclosed() {
						return
					}
				}
			}

			if !oddclosed {
				oddele, oddok := <-oddstream
				if(oddok){fmt.Println(oddele)}
				if !oddok {
					closedstreams += 1
					oddclosed = true
					if allclosed() {
						return
					}
				}
			}
		}
	}()

	wg.Wait()
}
