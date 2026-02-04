package main

import "errors"

// Stats returns count, sum, and average of a slice of float64.
// Named returns make it clear what each value means.
func Stats(numbers []float64) (count int, sum float64, avg float64, err error) {
	if len(numbers) == 0 {
		err = errors.New("empty slice")
		return // naked return — all named values returned at their current state
	}

	// Using explicit returns here. I prefer explicit returns in most cases —
	// naked returns are fine for early error exits where zero values are correct,
	// but in the "happy path" with real computed values, explicit is clearer
	// about what's actually being returned.
	count = len(numbers)
	for _, n := range numbers {
		sum += n
	}
	avg = sum / float64(count)
	return count, sum, avg, nil
}

// SummarizeGrades takes student grades and returns pass/fail counts.
// A passing grade is >= 60.
func SummarizeGrades(grades map[string]int) (passed, failed int) {
	for _, grade := range grades {
		if grade >= 60 {
			passed++
		} else {
			failed++
		}
	}
	return passed, failed
}
