package main

import (
	"Udemy/Golang/47-access-modifier/helper"
	"fmt"
)

func main() {
	helper.SayHello("Eko")
	//helper.sayGoodBye("Eko") // akan error karena private
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) // akan error karena private
}
