package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

type Publisher struct{
	broker *Broker
}

func NewPublisher(b *Broker) *Publisher {
	return &Publisher{broker: b}
}

func(p Publisher) PublishEvents(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(){
		defer close(p.broker.stream)
		defer wg.Done()

		for i := range(1000){
		msg := json.RawMessage(fmt.Sprintf(`{"id":%d}`, i))
		
		select{
		case p.broker.stream <- Stream{message: msg, topic: "message"}:
		case <-ctx.Done():
			return 
		}
	}
	}()
}