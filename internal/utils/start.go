package utils

import "context"

type Startable interface {
	Start(context.Context) error
}

func StartProcess(ctx context.Context, startable Startable) <-chan error {
	errCh := make(chan error, 1)

	go func(ctx context.Context, startable Startable, ch chan<- error) {
		errCh <- startable.Start(ctx)
	}(ctx, startable, errCh)

	return errCh
}
