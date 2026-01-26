package main

import "fmt"

const (
	Factor1 = 5
	Factor2 = 10
)

func OptimalVariables() {

	sum1 := scaleValue(Factor1, 2)
	fmt.Println("Sum:", sum1)

	sum2 := scaleValue(Factor2, 3)
	fmt.Println("Sum:", sum2)

	total := sum1 + sum2
	fmt.Println("Total:", total)
}

func scaleValue(count int, factor int) int {
	return count * factor
}
