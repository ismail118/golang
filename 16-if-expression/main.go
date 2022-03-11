package main

import "fmt"

func main() {
	name := "Eko"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	}

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	// if dengan short statement
	if lenght := len(name); lenght > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
