package main

import "context"

const MAX_CHUNKS = 5

func IngestionWithContext(ctx context.Context, maxnum int) []<-chan int {
	// stream initilization
	streams := make([]<-chan int,0)

	// streaming numbers parallely
	numChunks := min(maxnum, MAX_CHUNKS)
	if(numChunks == 0) { return streams }

	streamchunks := maxnum / numChunks

	for k := range numChunks {
		start := k * streamchunks
		end := start + streamchunks
		if k == numChunks-1 {
			end = maxnum
		}
		stream := make(chan int,5)
		go func(st chan int, start int, end int) {
			defer close(st)
			for i := start; i < end; i++ {
				select {
				case st <- i:
				case <-ctx.Done():
					return
				}
			}
		}(stream, start, end)
		streams = append(streams, stream)
	}

	// returning streams
	return streams
}
