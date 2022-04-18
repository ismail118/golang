package _94_streaming_encoder

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

func TestStreamingEncoder(t *testing.T) {
	writer, _ := os.Create("./sample_output.json")
	encode := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Edo",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
		Age:        22,
		Married:    false,
	}

	err := encode.Encode(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
