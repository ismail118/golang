package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
	for more
	see documentation
	doc https://pkg.go.dev/container/ring
	ex https://pkg.go.dev/container/ring#pkg-examples
*/
func main() {
	var data *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// gunakan function ini untuk iterasi data ring jangan gunakan for loop ( akan kena invinite loop )
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
