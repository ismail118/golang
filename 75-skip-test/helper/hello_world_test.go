package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

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
	result := HelloWorld("One")
	if result != "Hello Two" {
		// unit test failed
		t.Fatal("Result must be Hello Two")
	}

	fmt.Println("TestHelloWorldTwo Done")
}
