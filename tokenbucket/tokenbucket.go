package main

import (
	"context"
	"sync"
	"time"
)

type RateLimiter interface{
	Allow() bool
	triggerAutoFill(context.Context)
}

const RATE_IN_S = 10 // in seconds
const NO_OF_ALLOWS = 12 // no of allows

type TokenBucket struct{
	bucket int
	r int
	lastFilledAt time.Time
	capacity int // max capacity
}

var tb RateLimiter = nil
var mutex sync.Mutex
var lk chan int = make(chan int,1)

func NewDefaultRateLimiter(ctx context.Context, c int, r int) RateLimiter {
	mutex.Lock()
	defer mutex.Unlock()

	if(tb != nil){
		return tb
	}

	tb = &TokenBucket{
		bucket: c,
		r: r,
		lastFilledAt: time.Now(),
		capacity: c,
	}
	tb.triggerAutoFill(ctx)

	return tb
}

func (tb *TokenBucket) Allow() bool{
	lk<-1
	defer func(){<-lk}()

	if(tb.bucket >= 1) {
		tb.bucket-- 
		return true
	}
	return false
}

func (tb *TokenBucket) triggerAutoFill(ctx context.Context) {
	tk := time.NewTicker(time.Second * 1)
	go func(){
		for {
			select{
			case <-tk.C:
				lk <- 1
				if(tb.capacity != tb.bucket) {tb.fill()}
				<-lk
			case <-ctx.Done():
				tk.Stop()
				return
			}
		}
	}()
}

func (tb *TokenBucket) fill(){
	tb.bucket = min(tb.bucket + tb.r, tb.capacity)
	tb.lastFilledAt = time.Now()
}