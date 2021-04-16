package main

func main() {
	cache := map[int]int{
		0: 0,
		1: 1,
	}
	fib(10, cache)
}

func fib(i int, cache map[int]int) int {
	if v, ok := cache[i]; ok {
		return v
	}
	cache[i] = fib(i-1, cache) + fib(i-2, cache)
	return cache[i]
}
