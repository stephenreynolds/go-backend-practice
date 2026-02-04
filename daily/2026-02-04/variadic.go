package main

import (
	"fmt"
	"strings"
)

// Sum takes any number of ints and returns their total.
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// JoinWith joins strings using a separator.
// First arg is the separator, rest are the strings to join.
func JoinWith(sep string, parts ...string) string {
	return strings.Join(parts, sep)
}

// WrapErrors combines multiple errors into one.
// Skips nil errors. Returns nil if all are nil.
func WrapErrors(errs ...error) error {
	var msgs []string
	for _, e := range errs {
		if e != nil {
			msgs = append(msgs, e.Error())
		}
	}
	if len(msgs) == 0 {
		return nil
	}
	return fmt.Errorf("multiple errors: %s", strings.Join(msgs, "; "))
}
