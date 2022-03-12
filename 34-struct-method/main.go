package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func main() {
	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     30,
	}

	joko.sayHello()
}
