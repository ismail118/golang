package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// kode program pass by value
	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 Address = address1

	address2.City = "Bandung"

	fmt.Println(address1) // address1 tidak berubah
	fmt.Println(address2)

	// kode program pass by reference
	// menggunakan operator &
	var address3 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address4 *Address = &address3

	address4.City = "Depok"

	fmt.Println(address3)
	fmt.Println(address4)

	// menggunakan operator * (di gunakan untuk merubah semua data address5)
	var address5 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address6 *Address = &address5

	*address6 = Address{"Malang", "Jawa Barat", "Indonesia"}

	fmt.Println(address5)
	fmt.Println(address6)

	// function new
	var alamat1 *Address = new(Address)
	alamat2 := alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
