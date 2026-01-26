package main

import "fmt"

// Global variable
var total int

func variables() {
	// local variable
	var sum int
	{
		// Block level local variable
		count := 5
		sum = scaleValue(&count, 2)
		fmt.Println("Sum:", sum)
	}
	total = sum
	{
		count := 10
		sum = scaleValue(&count, 3)
		fmt.Println("Sum:", sum)
	}
	total := total + sum
	fmt.Println("Total:", total)
}

func scaleValue(count *int, factor int) int {
	// Accessing local variable
	sum := *count * factor
	return sum
}
