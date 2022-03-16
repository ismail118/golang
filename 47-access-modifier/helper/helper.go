package helper

import "fmt"

var version = "1.0.0" // tidak bisa diakses dari luar
var Application = "golang"

// Tidak bisa diakses dari luar package
func sayGoodBye(name string) {
	fmt.Println("Good Bye ", name)
}

func SayHello(name string) {
	fmt.Println("Hello ", name)
}
