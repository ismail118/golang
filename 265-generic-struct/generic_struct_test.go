package _65_generic_struct

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Eko",
		Second: "Khannedy",
	}

	fmt.Println(data)
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Eko",
		Second: "Khannedy",
	}

	assert.Equal(t, "Budi", data.ChangeFirst("Budi"))
	assert.Equal(t, "Hello Eko", data.SayHello("Eko"))
}
