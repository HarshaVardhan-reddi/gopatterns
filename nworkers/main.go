package main

import (
	"fmt"
	"sync"
)

const MAX_JOBS = 10

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main(){
	fmt.Println("Process jobs from a long running go routines")
	stream := make(chan int)

	// sending 1000 jobs to be processed by n workers
	go func(){
		for i := range(1000000){
		stream <- i
	}
	close(stream)
	}()

	for range(MAX_JOBS) {
		wg.Add(1)
		go processJobs(stream)
	}
	wg.Wait()
}