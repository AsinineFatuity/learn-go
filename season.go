package main

const JANUARY = 1
const FEBRUARY = 2
const MARCH = 3
const APRIL = 4
const MAY = 5
const JUNE = 6
const JULY = 7
const AUGUST = 8
const SEPTEMBER = 9
const OCTOBER = 10
const NOVEMEBER = 11
const DECEMBER = 12

func Season(month int) string {
	switch month {
	case 1, 2, 12:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	default:
		return "Season unknown"
	}
}
