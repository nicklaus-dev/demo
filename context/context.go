package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printHello(ctx); err != nil {
			fmt.Println(err)
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printCallback(ctx); err != nil {
			fmt.Println(err)
		}
	}()

	wg.Wait()
}

func do(ctx context.Context) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(time.Minute):
		return "Done", nil
	}
}

func printHello(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	s, err := do(ctx)
	if err != nil {
		return err
	}
	if s == "Done" {
		fmt.Println("Hello World!")
	}
	return nil
}

func printCallback(ctx context.Context) error {
	s, err := do(ctx)
	if err != nil {
		return err
	}
	if s == "Done" {
		fmt.Println("Hello World Callback!")
	}
	return nil
}
