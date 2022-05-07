package _66_generic_interface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func (m *MyData[T]) SetValue(value T) {
	m.Value = value
}

func TestGenericInterface(t *testing.T) {
	myData := &MyData[string]{}
	result := ChangeValue[string](myData, "Eko")

	assert.Equal(t, "Eko", result)
}
