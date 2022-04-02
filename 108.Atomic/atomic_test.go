package _08_Atomic

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	group := sync.WaitGroup{}
	var x int64 = 0

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				x += 1
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Hasil", x)
}
