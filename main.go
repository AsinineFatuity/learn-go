package main

import "fmt"

func PrintMessage(message string) (test string) {
	return message
}

func ConstantsExample() {
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
func main() {
	a := PrintMessage("Hi Eva, bride to be!")
	fmt.Println(a)
	ConstantsExample()
}
