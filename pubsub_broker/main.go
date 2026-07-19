package main

import (
	"context"
	"fmt"
	"sync"
)

func main(){
	fmt.Println("welcome to pubsub broker")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := &sync.WaitGroup{}
	// initializations
	subscriber := NewSubscriber("message processor")
	broker := NewBroker()
	publisher := NewPublisher(broker)
	broker.Subscribe("message",*subscriber)
	// processing
	publisher.PublishEvents(ctx,wg)
	broker.ProcessEvents(ctx, wg)
	subscriber.ConsumeEvents(ctx,wg)
	wg.Wait()
}