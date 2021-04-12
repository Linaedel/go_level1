package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(containsAllowedSymbols(6, 12, 3, 7))
}

func containsAllowedSymbols(numbers ...int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)
	sort.Ints(result) // Или собственная реализация одного из алгоритмов сортировки
	return result
}
