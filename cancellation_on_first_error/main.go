package main

import (
	"context"
	"fmt"
)

func main(){
	fmt.Println("cancellation on the first error")
	ctx, cancel := context.WithCancel(context.Background())
	errpipe := RunWorkersWithContext(ctx, 100000, false)
	err := MonitorWorkers(ctx,errpipe,cancel)
	if err != nil{
		fmt.Println("Workers cancelled because of error", err)
	}else{
		fmt.Println("Congratulations! No error")
	}
}