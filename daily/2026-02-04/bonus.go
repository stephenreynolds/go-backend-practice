package main

import "fmt"

// Retry calls fn up to N times until it succeeds.
// Returns the last error if all attempts fail.
func Retry(attempts int, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		fmt.Printf("  attempt %d failed: %v\n", i+1, err)
	}
	return fmt.Errorf("all %d attempts failed, last error: %w", attempts, err)
}

// Filter returns a new slice containing only elements where predicate returns true.
func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, n := range nums {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}

// Reduce folds a slice down to a single value using fn, starting from initial.
func Reduce(nums []int, initial int, fn func(int, int) int) int {
	acc := initial
	for _, n := range nums {
		acc = fn(acc, n)
	}
	return acc
}

// Memoize wraps a function so repeated calls with the same argument
// return cached results instead of recomputing.
func Memoize(fn func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(x int) int {
		if val, ok := cache[x]; ok {
			fmt.Printf("  cache hit for %d\n", x)
			return val
		}
		result := fn(x)
		cache[x] = result
		return result
	}
}
