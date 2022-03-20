package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// unit test failed
		panic("Result is not Hello Eko")
	}
}
