package main

import "fmt"

func PrintMessage(message string) (test string) {
	fmt.Println(message)
	return message
}

func ConstantsExample() {
	const PI float64 = 3.14
	var n int = 6
	fmt.Println(float64(n+5) * PI)
}
func main() {
	a := PrintMessage("Hi Eva, bride to be!")
	fmt.Println(a)
	ConstantsExample()
}
