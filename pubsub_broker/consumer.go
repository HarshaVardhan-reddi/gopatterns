package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

func ConsumeEvents(stream <-chan json.RawMessage) {

	wg := sync.WaitGroup{}

	wg.Go(func() {
		for {
			msg, ok := <-stream
			if !ok {
				return
			}
			time.Sleep(1 * time.Second)
			fmt.Println("Processed ", string(msg))
		}
	})

	wg.Wait()
}
