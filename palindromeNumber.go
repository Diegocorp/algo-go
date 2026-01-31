package main

import (
	"fmt"
)

func main() {
	number := 121

	fmt.Print(palindromeNumber(number))
}

func palindromeNumber(number int) bool {
	originalMap := make(map[int]int)
	// finalMap := make(map[int]int)

	for index, digit := range number {
		originalMap[digit] = index
	}

	return originalMap
