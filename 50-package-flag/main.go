package main

import (
	"flag"
	"fmt"
)

/*
	for more
	read documentation
	doc https://pkg.go.dev/flag
	ex https://pkg.go.dev/flag#pkg-examples
*/

func main() {
	var host *string = flag.String("host", "localhost", "Put your host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "Put your number")

	flag.Parse() // penting untuk di panggil sebelum menggunakan variable

	fmt.Println(*host, *user, *password, *number)
}
