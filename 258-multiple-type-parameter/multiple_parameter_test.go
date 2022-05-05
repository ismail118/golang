package _58_multiple_type_parameter

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1, param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("Eko", 100)
	MultipleParameter[int, string](100, "Eko")
}
