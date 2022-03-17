package main

import (
	"fmt"
	"sort"
)

/*
	for more
	see documentation
	doc https://pkg.go.dev/sort
	ex https://pkg.go.dev/sort#pkg-examples
*/

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Eko", 30},
		{"Budi", 35},
		{"Joko", 25},
		{"Adit", 23},
	}

	fmt.Println(users)

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
