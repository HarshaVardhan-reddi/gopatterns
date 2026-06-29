package main

import (
	"context"
	"fmt"
	"net/http"

	"time"
)

var mutex chan struct{} = make(chan struct{}, 10)

func process(url string) {

	mutex <- struct{}{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

	defer cancel()
	defer func() { <-mutex }()
	defer wg.Done()

	fmt.Println("executing",url)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if(err != nil){
		fmt.Println("error here", err.Error())
		return 
	}
	defer resp.Body.Close()
}
