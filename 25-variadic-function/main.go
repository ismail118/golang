package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	// slice paramter for varadic function
	numbers := []int{1, 2, 3, 4, 5}
	totalSlice := sumAll(numbers...)
	fmt.Println(totalSlice)
}
