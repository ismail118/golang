package main

import (
	"fmt"
	"strings"
)

/*
	for more
	see documentation
	doc https://pkg.go.dev/strings
	ex https://pkg.go.dev/strings#pkg-examples
*/

func main() {
	// contains
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Contains("Eko Kurniawan", "Budi"))

	// split
	fmt.Println(strings.Split("Eko Kurniawan Khannedy", " "))

	// toLowwer
	fmt.Println(strings.ToLower("Eko Kurniawan Khannedy"))
	// toUpper
	fmt.Println(strings.ToUpper("Eko Kurniawan Khannedy"))

	// trim
	fmt.Println(strings.Trim("            Eko Kurniawan          ", " "))

	// replaceAll
	fmt.Println(strings.ReplaceAll("Eko Eko Eko", "Eko", "Budi"))
}
