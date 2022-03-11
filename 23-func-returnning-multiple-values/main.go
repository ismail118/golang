package main

import "fmt"

func getFullName() (string, string) {
	return "Eko", "Khannedy"
}

func main() {
	firstName, lastNmae := getFullName()
	fmt.Println(firstName, lastNmae)

	// mehghiraukan return value
	firstName1, _ := getFullName()
	fmt.Println(firstName1)
}
