package main

import (
	_ "Udemy/Golang/48-package-initialization/blank_identifier" // contoh penggunaan blank identifier
	"Udemy/Golang/48-package-initialization/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
