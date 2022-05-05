package _57_type_parameter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	assert.NotNil(t, true)
}

func TestLength(t *testing.T) {
	var resultString string = Length[string]("Eko")
	assert.Equal(t, "Eko", resultString)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}
