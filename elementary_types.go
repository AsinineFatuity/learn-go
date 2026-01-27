package main

import "fmt"

func ElementaryTypes() {
	var num float32 = 3.456
	// represent number with a precision of m digits and width of n digits %n.mg
	fmt.Printf("Number is %2.4e\n", num)
	// width in non digit values
	lang := "Go"
	fmt.Printf("|%5s|\n", lang)
	fmt.Printf("|%-5s|\n", lang)

}
