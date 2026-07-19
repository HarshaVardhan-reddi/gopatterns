package main

import (
	"context"
	"fmt"
	"sync"
)

type Subscriber struct {
	name   string
	stream chan Stream
}

func NewSubscriber(name string) *Subscriber {
	return &Subscriber{name: name, stream: make(chan Stream)}
}

func (s Subscriber) ConsumeEvents(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case st, ok := <-s.stream:
				if !ok {
					return
				}
				fmt.Println(s.name, " is processing event", string(st.message))
			case <-ctx.Done():
				return
			}
		}
	}()
}
