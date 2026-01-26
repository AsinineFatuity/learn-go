package main

import "fmt"

func printMessage(message string) (test string) {
	return message
}

func constantsExample() {
	const PI float64 = 3.14
	var n int = 6
	fmt.Println(float64(n+5) * PI)
	// enumeration example
	type LanguageKnown int
	const (
		PYTHON LanguageKnown = iota
		TYPESCRIPT
		GOLANG
	)
	fmt.Println(PYTHON, TYPESCRIPT, GOLANG)
}

func variablesExample() {
	/*
		variables are named in camel case (exampleVariable) but if they are to be exported the must
		 must start with capital letter (ExampleVariable)
	*/
	var num = 10
	fmt.Printf("The number is %d\n", num)
	var decision = true
	fmt.Printf("The decision is %t\n", decision)
}
func main() {
	a := printMessage("Hello Go World!")
	fmt.Println(a)
	// constantsExample()
	// variablesExample()
	OptimalVariables()
}
