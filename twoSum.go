package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Print(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)

	for i, num := range nums {
		missing := target - num

		index, ok := myMap[missing]

		if ok {
			return []int{index, i}
		}

		myMap[num] = i
	}

	return []int{}
}
