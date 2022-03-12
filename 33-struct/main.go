package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 32

	fmt.Println(eko)

	// struct literals
	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	// tidak recomended menggunakan cara yg di bawah ini
	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)
}
