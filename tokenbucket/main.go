// 5 Token-bucket rate limiter
// Build a rate limiter based on the token-bucket model. It is configured with a capacity B (the maximum
// burst size) and a refill rate R (tokens added per second, up to the cap). Expose a non-blocking Allow()
// bool: each call tries to consume one token and returns true if a token was available, or false
// immediately if not. Tokens accrue over time so an idle caller builds up a burst allowance, but the bucket
// never holds more than B tokens. Must be safe to call from many goroutines at once.
// REQUIREMENTS
// – Allow() consumes one token and never blocks — returns true/false right away.
// – Refill continuously over real time, capped at B; idle time should not over-fill.
// – Concurrency-safe under heavy parallel calls.
// – Consider lazy refill (elapsed time per call) versus a background ticker; defend your choice

package main

import (
	"context"
	"fmt"
	"sync"
)

const MAX_GO_ROUTINES = 100000


func main(){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	rt := NewDefaultRateLimiter(ctx,12,10)

	for range MAX_GO_ROUTINES {
		wg.Add(1)
		go func (){
			defer wg.Done()
			fmt.Println(rt.Allow())
		}()
	}

	wg.Wait()
}