package _93_streaming_decoder

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestStreamingDecoder(t *testing.T) {
	reader, _ := os.Open("sample.json")
	decoder := json.NewDecoder(reader)

	customer := Customer{}
	_ = decoder.Decode(&customer)

	fmt.Println(customer)
}
