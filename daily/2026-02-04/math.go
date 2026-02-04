package main

import "errors"

// Divide returns quotient and remainder.
// Returns an error if divisor is zero.
func Divide(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	return a / b, a % b, nil
}

// MinMax returns the minimum and maximum of a slice.
// Returns an error if the slice is empty.
func MinMax(nums []int) (int, int, error) {
	if len(nums) == 0 {
		return 0, 0, errors.New("empty slice")
	}

	min, max := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max, nil
}

// Swap returns its two arguments in reverse order.
func Swap(a, b string) (string, string) {
	return b, a
}
