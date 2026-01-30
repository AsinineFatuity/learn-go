package main

import (
	"fmt"
)

// Factorial calculates the factorial of a non-negative integer n.
// It returns 0 as the factorial value if n is negative.
func Factorial(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1 // Factorial of 0 is 1
	}

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	return factorial
}

func ExecuteFactorial() {
	n := 10
	x := -10

	result := Factorial(n)
	fmt.Printf("Factorial of %d is %d\n", n, result)

	result = Factorial(x)
	fmt.Printf("Factorial of %d is %d\n", x, result)
}
