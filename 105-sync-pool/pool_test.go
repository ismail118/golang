package _05_sync_pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool sync.Pool

	pool.Put("Eko")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}

func TestPoolV2(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Eko")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
