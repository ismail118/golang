package main

import "fmt"

func endApp() {
	fmt.Println("Applikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Applikasi Error")
	}
	fmt.Println("Applikasi Berjalan")
}

func main() {
	//runApp(false)

	// triger panic
	runApp(true)
}
