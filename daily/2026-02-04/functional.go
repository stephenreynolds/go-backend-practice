package main

// Apply takes a slice of ints and a transform function,
// returns a new slice with the function applied to each element.
func Apply(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}

// MakeMultiplier returns a function that multiplies by n.
// This is a closure — it captures n.
func MakeMultiplier(n int) func(int) int {
	return func(x int) int {
		return x * n
	}
}

// Compose takes two functions and returns their composition.
// Compose(f, g)(x) = f(g(x))
func Compose(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}
