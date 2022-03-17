package main

import (
	"fmt"
	"regexp"
)

/*
	for more
	see documentation
	doc regex syntax https://github.com/google/re2/wiki/Syntax
	doc https://pkg.go.dev/regexp
	ex https://pkg.go.dev/regexp#pkg-examples
*/
func main() {
	var regex = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eKo"))

	fmt.Println(regex.FindAllString("eto eko edo egi elo ego", 10))
}
