package main

import (
	"context"
	"fmt"
)

func main(){
	fmt.Println("Merge k channels pattern")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	numstreams := IngestionWithContext(ctx,1000)
	junction := MergeStreamsWithContext(ctx,numstreams...)
	ProcessMergedStreamWithContext(ctx,junction)
}