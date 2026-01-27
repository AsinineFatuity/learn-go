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

func Bitwise() {
	x := 3 //0011
	y := 5 //0101
	z := (x ^ y) & (x << 2)
	// x ^ y = 0011 ^ 0101 = 0110
	// x << 2 = 0011 << 2 = 1100
	// 0110 & 1100 = 0100 = 4
	fmt.Println("Bitwise answer", z)
}
