package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	//resultInt := result.(int) // will produce error
	//fmt.Println(resultInt)

	// type assertions using switch
	result1 := random()
	switch value := result1.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}
