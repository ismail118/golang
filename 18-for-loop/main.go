package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// for dengan statement (init statement dan post statement)
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	// for range
	slice := []string{"Eko", "Kurniawan", "Khannedy"}

	for i, name := range slice {
		fmt.Printf("%d.%s \n", i, name)
	}
	for _, name := range slice {
		fmt.Printf("%s \n", name)
	}

	mapPersone := make(map[string]string)
	mapPersone["name"] = "Eko"
	mapPersone["title"] = "Dev"

	for key, value := range mapPersone {
		fmt.Printf("%s : %s \n", key, value)
	}
}
