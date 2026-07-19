package main

import (
	"context"
	"encoding/json"
	"sync"
)

type Stream struct{
	topic string
	message json.RawMessage
}

type Broker struct{
	stream chan Stream
	topicfreq map[string]int
	topic_subs map[string][]Subscriber
}

func NewBroker() *Broker {
	return &Broker{
		stream: make(chan Stream),
		topic_subs: make(map[string][]Subscriber),
		topicfreq: make(map[string]int),
	}
}

func (b *Broker) ProcessEvents(ctx context.Context, wg *sync.WaitGroup){
	wg.Add(1)
	go func(){
		
		// subscriber termination
		terminate_sub := func(){
			for _, subs := range(b.topic_subs){
				for _, sub := range(subs){
					close(sub.stream)
				}
			}
		}

		defer wg.Done()
		defer terminate_sub()

		for{
			select{
			case st, ok := <-b.stream:
				if(!ok){return}
				b.topicfreq[st.topic]++
				for _,subsciber := range(b.topic_subs[st.topic]) {
					select{
					case subsciber.stream <- st:
					case <-ctx.Done(): return 
					}
				}
			case <-ctx.Done(): return
			}
		}
	}()

}

func (b *Broker) Subscribe(topic string, s Subscriber){
	b.topic_subs[topic] = append(b.topic_subs[topic], s)
}