package _06_sync_map

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func addToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	var data *sync.Map
	var group *sync.WaitGroup

	for i := 0; i < 100; i++ {
		go addToMap(data, i, group)
	}

	group.Wait()

	time.Sleep(3 * time.Second)
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
