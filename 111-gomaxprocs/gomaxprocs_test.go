package _11_gomaxprocs

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println(totalGoroutine)
}
