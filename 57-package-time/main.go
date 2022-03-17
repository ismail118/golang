package main

import (
	"fmt"
	"time"
)

/*
	for more
	see documentation
	doc https://pkg.go.dev/time
	ex https://pkg.go.dev/time#pkg-examples
*/
func main() {
	// mendapatkan waktu sekarang
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// membuat tanggal
	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2022-03-17")
	fmt.Println(parse)
}
