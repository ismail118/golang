package _19_context_with_timeout

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++

				time.Sleep(1 * time.Second) // untuk simulasi time out
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("TOTAL GOROUTINE", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 1*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("TOTAL GOROUTINE", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)
	fmt.Println("TOTAL GOROUTINE", runtime.NumGoroutine())
}
