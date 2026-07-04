package main

import (
	"context"
)

func MonitorWorkers(ctx context.Context, errchan <-chan error, cancel context.CancelFunc) error {
	defer cancel()
	var err error
	for er := range errchan {
		if err == nil {
			err = er
			cancel()
		}
	}
	return err
}