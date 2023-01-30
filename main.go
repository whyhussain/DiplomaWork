package main

import (
	app "DiplomaWork/internal/app"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)

	go func() {
		sigCH := make(chan os.Signal)
		signal.Notify(sigCH, syscall.SIGTERM, syscall.SIGINT)
		errCh <- fmt.Errorf("%s", <-sigCH)
	}()

	app.StartNumbleServie(ctx, errCh)

}
