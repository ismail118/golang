package main

import "fmt"

type Address struct {
	Country string
}

func ChangeAddressToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func ChangeAddressToIndonesia2(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"USA"}
	ChangeAddressToIndonesia(address)

	fmt.Println(address) // tidak berubah

	// menggunakan function 2
	ChangeAddressToIndonesia2(&address)

	fmt.Println(address) // berubah
}
