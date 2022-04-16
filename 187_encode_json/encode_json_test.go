package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestMarshal(t *testing.T) {
	LogJson("Eko")
	LogJson(1)
	LogJson(true)
	LogJson([]string{"Eko", "Edo", "Eto"})
}
