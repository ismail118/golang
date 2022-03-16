package main

import (
	"fmt"
	"os"
)

/*
	for more
	read documentation
	doc https://pkg.go.dev/os@go1.18
	ex https://pkg.go.dev/os@go1.18#pkg-examples
*/

func main() {
	// argument
	args := os.Args
	fmt.Println(args)

	// hostname
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err)
	}

	// get environment variable
	text := os.Getenv("GOPATH")
	fmt.Println(text)
}
