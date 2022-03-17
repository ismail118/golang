package main

import (
	"fmt"
	"reflect"
)

/*
	for more
	see documentation
	doc https://pkg.go.dev/reflect
	ex https://pkg.go.dev/reflect#pkg-examples
*/

type Sample struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("requried") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Eko"}
	var sampleType reflect.Type = reflect.TypeOf(sample)

	// ambil jumlah field
	fmt.Println("jumlah field : ", sampleType.NumField())

	// ambil name field ke 0
	fmt.Println("name field : ", sampleType.Field(0).Name)

	// ambil tag field ke 0
	fmt.Println("required : ", sampleType.Field(0).Tag.Get("required"))
	fmt.Println("required : ", sampleType.Field(0).Tag.Get("max"))
	fmt.Println("required : ", sampleType.Field(0).Tag.Get("min")) // jika tag tidak ada maka

	// membuat function falidasi menggunakan flag
	fmt.Println(IsValid(sample))
}
