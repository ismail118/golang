package _61_type_sets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min[T Number](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, int(100), Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](100, 200))
	assert.Equal(t, float64(100.0), Min[float64](100.0, 200.0))
}
