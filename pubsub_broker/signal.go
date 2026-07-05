package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"syscall"
)

func Signaling() context.Context {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	go func(){
		defer stop()
		for{
			<-ctx.Done()
			fmt.Println("***Stopping application gracefully as stop signal received***")
			return
		}
	}()
	return ctx
}