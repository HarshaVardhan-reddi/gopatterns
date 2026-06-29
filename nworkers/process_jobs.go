package main

import (
	"fmt"
)

func processJobs(stream <-chan int) {
	defer wg.Done()
	for{
		val, ok := <-stream
		if !ok {
			return
		}
		fmt.Println("processing val", val)
	}
}
