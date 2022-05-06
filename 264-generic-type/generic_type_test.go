package _64_generic_type

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBag(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 6}
	PrintBag[int](numbers)

	names := Bag[string]{"Eko", "Budi", "Joko"}
	fmt.Println(names)
	PrintBag[string](names)
}
