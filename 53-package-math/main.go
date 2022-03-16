package main

import (
	"fmt"
	"math"
)

/*
	for more
	see documentation
	doc https://pkg.go.dev/math
	ex https://pkg.go.dev/math#pkg-examples
*/

func main() {
	// membulatkan ke atas
	fmt.Println(math.Ceil(1.40))

	// membulatkan ke bawah
	fmt.Println(math.Floor(1.7))

	// membulatkan keatas atau ke bawah sesuai yg paling dekat
	fmt.Println(math.Round(1.7))
	fmt.Println(math.Round(1.3))

	// mecari terbesar
	fmt.Println(math.Max(1, 2))

	// mencari terkecil
	fmt.Println(math.Min(1, 2))
}
