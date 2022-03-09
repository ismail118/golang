package main

import "fmt"

func main() {
	const firstName string = "Eko"
	const lastName = "Khannedy"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// multiple deklarasi constant

	const (
		name     = "ismail"
		fullName = "ismail alfiyasin"
		age      = 24
	)

	fmt.Println(name, fullName, age)
}
