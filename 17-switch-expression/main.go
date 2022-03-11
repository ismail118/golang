package main

import "fmt"

func main() {
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan donk")
		fmt.Println("Kenalan donk")
	}

	// swhitch dengan short statement
	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("Nama terlalu panjang")
	case lenght > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
