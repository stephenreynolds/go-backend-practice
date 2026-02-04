package main

// ProcessScores takes raw test scores, filters out invalids (< 0 or > 100),
// applies a curve function, and returns stats.
func ProcessScores(scores []int, curve func(int) int) (count int, sum float64, avg float64, err error) {
	// Filter valid scores and apply curve
	var curved []float64
	for _, s := range scores {
		if s < 0 || s > 100 {
			continue
		}
		curved = append(curved, float64(curve(s)))
	}

	return Stats(curved)
}
