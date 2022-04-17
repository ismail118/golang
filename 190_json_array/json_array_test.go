package _90_json_array

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
		Age:        0,
		Married:    false,
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"Kurniawan","LastName":"Khannedy","Age":0,"Married":false,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}
	json.Unmarshal(jsonBytes, &customer)

	fmt.Println(customer)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Ode",
		Addresses: []Address{
			{
				Street:     "jalan",
				Country:    "depok",
				PostalCode: "123",
			},
			{
				Street:     "jalan-jalan",
				Country:    "depok",
				PostalCode: "123333",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Ode","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"jalan","Country":"depok","PostalCode":"123"},{"Street":"jalan-jalan","Country":"depok","PostalCode":"123333"}]}`
	jsonByte := []byte(jsonString)

	customer := Customer{}
	_ = json.Unmarshal(jsonByte, &customer)

	fmt.Println(customer)
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"jalan","Country":"depok","PostalCode":"123"},{"Street":"jalan-jalan","Country":"depok","PostalCode":"123333"}]`
	jsonByte := []byte(jsonString)

	addresses := []Address{}
	_ = json.Unmarshal(jsonByte, &addresses)

	fmt.Println(addresses)
}

func TestOnlyJsonArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "jalan",
			Country:    "depok",
			PostalCode: "123",
		},
		{
			Street:     "jalan-jalan",
			Country:    "depok",
			PostalCode: "123333",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
