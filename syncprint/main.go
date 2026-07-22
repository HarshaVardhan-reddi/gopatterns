package main

import (
	"fmt"
	
	"sync"
)

func main(){
	var wg sync.WaitGroup

	sharedsig1 := make(chan struct{},1)
	sharedsig2 := make(chan struct{},1)

	hello := hellostream(&wg, sharedsig1, sharedsig2)
	hi := histream(&wg, sharedsig1, sharedsig2)

	// kickoff: unblocks histream so "hi" prints first
	sharedsig2 <- struct{}{}

	// single consumer reads in the signal-determined order: hi, then hello
	streams := []<-chan string{hi, hello}
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		for range(5){
			for _, ch := range(streams){
				fmt.Println(<-ch)
			}
		}
	}()

	wg.Wait()
}

func hellostream(wg *sync.WaitGroup, sig1 chan struct{}, sig2 chan struct{}) <-chan string{
	stream := make(chan string)
	wg.Add(1)

	go func(){
		defer wg.Done()
		defer close(stream)
		for range(5){
			<-sig1
			stream <- "hello"
			sig2<-struct{}{}
		}
	}()

	return stream
}

func histream(wg *sync.WaitGroup, sig1 chan struct{}, sig2 chan struct{}) <-chan string{
	stream := make(chan string)
	wg.Add(1)
	
	go func(){
		defer wg.Done()
		defer close(stream)
		for range(5){
			<-sig2
			stream <- "hi"
			sig1<-struct{}{}
		}
	}()

	return stream
}