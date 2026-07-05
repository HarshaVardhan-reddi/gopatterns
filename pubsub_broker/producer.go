package main

import (
	"context"
	"encoding/json"
)

func ProduceEventsWithContext(ctx context.Context) chan json.RawMessage {

	events := make([]json.RawMessage,0)
	stream := make(chan json.RawMessage,10)

	for i := range(1000000){
		event := map[string]int{"_id": i}
		
		rawmsg, err := json.Marshal(event)
		if(err !=nil){
			panic(err)
		}

		events = append(events, rawmsg)
	}

	go func(){
		defer close(stream)
		for _, event := range(events){
			select{
			case stream <- event:
			case <-ctx.Done():
				return
			}
		}
	}()

	return stream
}