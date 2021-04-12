package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsAllowedSymbols(6, 12, 3, 7))
}

func containsAllowedSymbols(numbers ...int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)
	for i := 1; i < len(result); i++ {
		temp := result[i]
		index := i
		for j := i - 1; j > -1; j-- {
			if result[j] < temp {
				break
			}
			result[j+1] = result[j]
			index = j
		}
		result[index] = temp
	}
	return result
}
