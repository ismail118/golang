package main

import "fmt"

func main() {
	// 1
	var name string

	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	// 2
	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 38
	fmt.Println(age)

	// 3
	country := "indonesia"
	fmt.Println(country)

	//4 deklarasi multipe variable
	var (
		firstName = "ismail"
		lastName  = "alfiyasin"
	)

	fmt.Println(firstName, lastName)
}
