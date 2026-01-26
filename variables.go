package main

import "fmt"

func OptimalVariables() {
	// local variable
	var sum int
	{
		// Block level local variable
		count := 5
		sum = scaleValue(&count, 2)
		fmt.Println("Sum:", sum)
	}
	total := sum
	{
		count := 10
		sum = scaleValue(&count, 3)
		fmt.Println("Sum:", sum)
	}
	total += sum
	fmt.Println("Total:", total)
}

func scaleValue(count *int, factor int) int {
	// Accessing local variable
	scaledValue := *count * factor
	return scaledValue
}
