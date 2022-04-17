package _89_decode_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
}

func TestDecodeJson(t *testing.T) {
	jsonRequest := `{"FirstName":"Eko", "MiddleName":"Kurniawan", "LastName":"Khannedy"}`
	jsonBytes := []byte(jsonRequest)

	customer := Customer{}
	json.Unmarshal(jsonBytes, &customer)

	fmt.Println(customer)
}
