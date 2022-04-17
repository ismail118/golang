package _92_map

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapEncode(t *testing.T) {
	jsonRequest := `{"id":"1234","name":"Apple iphone","price":200000000}`
	jsonBytes := []byte(jsonRequest)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapDecode(t *testing.T) {
	products := map[string]interface{}{
		"id":    "1234",
		"name":  "Apple iphone",
		"price": 200000000,
	}

	bytes, _ := json.Marshal(products)
	fmt.Println(string(bytes))
}
