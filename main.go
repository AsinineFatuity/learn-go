package main

import "fmt"

func PrintMessage(message string) (test string) {
	fmt.Println(message)
	return message
}
func main() {
	var a string = PrintMessage("Hi Eva, bride to be!")
	fmt.Println(a)
}
