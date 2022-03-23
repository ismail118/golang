package _9_membuat_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go HelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}
