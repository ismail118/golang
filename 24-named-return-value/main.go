package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"
	return
}
func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)

}
