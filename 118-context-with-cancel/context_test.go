package _18_context_with_cancel

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
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println(runtime.NumGoroutine())
	cancel() // mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}
