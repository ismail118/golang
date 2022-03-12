package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi Error ", message)
	}
	fmt.Println("Applikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Applikasi Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("Eko")
}
