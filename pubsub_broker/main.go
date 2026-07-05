package main

import "fmt"

func main(){
	fmt.Println("Pub sub pattern in golang")
	ctx := Signaling()
	stream := ProduceEventsWithContext(ctx)
	ConsumeEvents(stream)
}