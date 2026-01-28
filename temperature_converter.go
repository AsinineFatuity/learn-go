package main

import "fmt"

type Celsius float32
type Fahrenheit float32

const DIVISOR Celsius = 9.0 / 5.0
const MARGIN Celsius = 32

func ToFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit
	temp = Fahrenheit((t * DIVISOR) + MARGIN)
	fmt.Println("Temperature in celsius:", t, "In Farenheit is: ", temp)
	return temp
}
