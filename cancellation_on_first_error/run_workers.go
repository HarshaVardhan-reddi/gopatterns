package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

func RunWorkersWithContext(ctx context.Context,n int, isSuccess bool) <-chan error {
	var wg *sync.WaitGroup = &sync.WaitGroup{}

	errpipe := make(chan error)
	maxloops := min(n,5);
	chunks := n / maxloops
	start := 0
	doWork := func(x int) error {
		if(x >= 0){
			fmt.Println("printing the number which is passed", x)
			return nil
		}else{
			return errors.New("number cannot be less than zero or greter than zero")
		}
	}

	for range(maxloops){
		end := start + chunks
		wg.Add(1)
		go func(s int, e int){
			defer wg.Done()
			for i := s; i < e; i++{
				select{
				case <-ctx.Done():
					return
				default:
					err := doWork(i)
					if err != nil{
						errpipe <- err
						return
					}
				}
			}
		}(start, end)
		start = end
	}

	if !isSuccess{
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			select{
			case <-ctx.Done():
				return 
			default:
				err := doWork(-1)
				errpipe <- err
			}
		}()
	}

	go func ()  {
		defer close(errpipe)
		wg.Wait()
	}()

	return errpipe
}