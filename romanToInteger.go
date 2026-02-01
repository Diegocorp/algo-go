package main

import (
	"fmt"
)

func main() {
	romanNumber := "LVIII"
	fmt.Println(romanToInteger(romanNumber))
}

func romanToInteger(romanNumber string) int {
	integerNumber := 0

	romanNumeral := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < len(romanNumber); i++ {

		if i+1 < len(romanNumber) && romanNumeral[string(romanNumber[i])] < romanNumeral[string(romanNumber[i+1])] {
			integerNumber -= romanNumeral[string(romanNumber[i])]
		} else {
			integerNumber += romanNumeral[string(romanNumber[i])]
		}
	}
	return integerNumber
}
