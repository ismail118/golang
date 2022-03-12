package main

import "fmt"

func main() {
	name := "Eko"
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		name = "Budi"
		counter++
	}

	increment()
	increment()

	fmt.Println(name)
	fmt.Println(counter)
}
