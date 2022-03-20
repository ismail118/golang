package helper

import (
	"fmt"
	"testing"
)

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
	result := HelloWorld("OneFailTest")
	if result != "Hello One" {
		// unit test failed
		t.Error("Result must be Hello One")
	}

	fmt.Println("TestHelloWorldOne Done")
}

func TestHelloWorldTwo(t *testing.T) {
	result := HelloWorld("TwoFailTest")
	if result != "Hello Two" {
		// unit test failed
		t.Fatal("Result must be Hello Two")
	}

	fmt.Println("TestHelloWorldTwo Done")
}
