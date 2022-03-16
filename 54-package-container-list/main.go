package main

import (
	"container/list"
	"fmt"
)

/*
	for more
	see documentation
	doc https://pkg.go.dev/container/list
	ex https://pkg.go.dev/container/list#pkg-examples
*/
func main() {
	data := list.New()
	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")

	// iterasi //

	// dari depan ke belakang
	for e := data.Front(); e != nil; e.Next() {
		fmt.Println(e.Value)
	}

	// dari belakang ke depan
	for e := data.Back(); e != nil; e.Prev() {
		fmt.Println(e.Value)
	}
}
