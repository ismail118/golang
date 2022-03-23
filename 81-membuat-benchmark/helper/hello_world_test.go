package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		HelloWorld("Eko")
	}
}

func BenchmarkHelloWorldEdo(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		HelloWorld("Edo")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Eko)",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "HelloWorld(Kurniawan)",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result)
	})
	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result)
	})
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Os Windows")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Fail resrult must be Hello Eko")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Fail result must be Hello Eko")
	fmt.Println("TestHelloWorldRequire Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Fail result must be Hello Eko")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// unit test failed
		t.Fail()
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")
	if result != "Hello Khannedy" {
		// unit test failed
		t.FailNow()
	}

	fmt.Println("TestHelloWorldKhannedy Done")
}

func TestHelloWorldOne(t *testing.T) {
	result := HelloWorld("One")
	if result != "Hello One" {
		// unit test failed
		t.Error("Result must be Hello One")
	}

	fmt.Println("TestHelloWorldOne Done")
}

func TestHelloWorldTwo(t *testing.T) {
	result := HelloWorld("Two")
	if result != "Hello Two" {
		// unit test failed
		t.Fatal("Result must be Hello Two")
	}

	fmt.Println("TestHelloWorldTwo Done")
}
