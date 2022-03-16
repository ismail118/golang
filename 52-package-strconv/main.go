package main

import (
	"fmt"
	"strconv"
)

/*
	for more
	see documentation
	doc https://pkg.go.dev/strconv
	ex https://pkg.go.dev/strconv#pkg-examples
*/

func main() {
	// string to bool
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	// string to int
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// int to string
	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	// int to string Itoa
	valueString := strconv.Itoa(1000)
	fmt.Println(valueString)

	// string to int Atoi
	valueInt, err := strconv.Atoi("1000")
	if err == nil {
		fmt.Println(valueInt)
	} else {
		fmt.Println(err.Error())
	}
}
