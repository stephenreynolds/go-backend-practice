package main

import (
	"errors"
	"fmt"
)

func main() {
	// ========== Part 1: Basic Functions & Multiple Returns ==========
	fmt.Println("=== Divide ===")
	q, r, err := Divide(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", q, r)

	_, _, err = Divide(17, 0)
	fmt.Printf("17 / 0 = error: %v\n", err)

	fmt.Println("\n=== MinMax ===")
	min, max, err := MinMax([]int{3, 1, 4, 1, 5, 9, 2, 6})
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("[3, 1, 4, 1, 5, 9, 2, 6] → min: %d, max: %d\n", min, max)
	}

	_, _, err = MinMax([]int{})
	fmt.Printf("[] → error: %v\n", err)

	fmt.Println("\n=== Swap ===")
	a, b := "hello", "world"
	fmt.Printf("Before: a=%s, b=%s\n", a, b)
	a, b = Swap(a, b)
	fmt.Printf("After:  a=%s, b=%s\n", a, b)

	// ========== Part 2: Named Returns ==========
	fmt.Println("\n=== Stats ===")
	numbers := []float64{85.5, 92.3, 78.1, 95.0, 88.7}
	count, sum, avg, err := Stats(numbers)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Numbers: %v\n", numbers)
		fmt.Printf("Count: %d, Sum: %.2f, Average: %.2f\n", count, sum, avg)
	}

	fmt.Println("\n=== Grades ===")
	grades := map[string]int{
		"Alice":   92,
		"Bob":     55,
		"Charlie": 78,
		"Diana":   43,
		"Eve":     88,
	}
	passed, failed := SummarizeGrades(grades)
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Passed: %d, Failed: %d\n", passed, failed)

	// ========== Part 3: Variadic Functions ==========
	fmt.Println("\n=== Variadic ===")
	fmt.Printf("Sum(1,2,3,4,5) = %d\n", Sum(1, 2, 3, 4, 5))

	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum(spread) = %d\n", Sum(nums...))

	fmt.Printf("JoinWith(\" | \", \"Go\", \"is\", \"fun\") = %q\n",
		JoinWith(" | ", "Go", "is", "fun"))

	err1 := errors.New("file not found")
	err2 := errors.New("permission denied")
	fmt.Printf("WrapErrors(nil, err1, nil, err2) = %q\n", WrapErrors(nil, err1, nil, err2))
	fmt.Printf("WrapErrors(nil, nil) = %v\n", WrapErrors(nil, nil))

	// ========== Part 4: Functions as Values ==========
	fmt.Println("\n=== Functions as Values ===")
	original := []int{1, 2, 3, 4, 5}
	double := func(x int) int { return x * 2 }
	square := func(x int) int { return x * x }

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Doubled:  %v\n", Apply(original, double))
	fmt.Printf("Squared:  %v\n", Apply(original, square))

	triple := MakeMultiplier(3)
	fmt.Printf("\ntriple := MakeMultiplier(3)\n")
	fmt.Printf("triple(7) = %d\n", triple(7))

	addOne := func(x int) int { return x + 1 }
	doubleThenAddOne := Compose(addOne, double)
	fmt.Printf("\ndoubleThenAddOne := Compose(addOne, double)\n")
	fmt.Printf("doubleThenAddOne(5) = %d\n", doubleThenAddOne(5))

	// ========== Part 5: Score Pipeline ==========
	fmt.Println("\n=== Score Pipeline ===")
	curve := func(score int) int {
		s := score + 5
		if s > 100 {
			return 100
		}
		return s
	}

	scores := []int{72, 85, -1, 91, 150, 68, 77, 95, 83}
	fmt.Printf("Raw scores: %v\n", scores)

	// Show the valid curved scores for clarity
	var validCurved []int
	for _, s := range scores {
		if s >= 0 && s <= 100 {
			validCurved = append(validCurved, curve(s))
		}
	}
	fmt.Printf("Valid scores after curve: %v\n", validCurved)

	count, sum, avg, err = ProcessScores(scores, curve)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Count: %d, Sum: %.2f, Average: %.2f\n", count, sum, avg)
	}

	// ========== Bonus Challenges ==========
	fmt.Println("\n=== Bonus: Retry ===")
	attempt := 0
	err = Retry(3, func() error {
		attempt++
		if attempt < 3 {
			return fmt.Errorf("connection refused")
		}
		fmt.Println("  connected!")
		return nil
	})
	fmt.Printf("Retry result: %v\n", err)

	fmt.Println("\n=== Bonus: Filter ===")
	evens := Filter([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Evens from [1..8]: %v\n", evens)

	fmt.Println("\n=== Bonus: Reduce ===")
	product := Reduce([]int{1, 2, 3, 4, 5}, 1, func(acc, n int) int {
		return acc * n
	})
	fmt.Printf("Product of [1..5]: %d\n", product)

	fmt.Println("\n=== Bonus: Memoize ===")
	slowSquare := func(x int) int {
		fmt.Printf("  computing %d²...\n", x)
		return x * x
	}
	memoSquare := Memoize(slowSquare)
	fmt.Printf("memoSquare(4) = %d\n", memoSquare(4))
	fmt.Printf("memoSquare(4) = %d\n", memoSquare(4)) // should hit cache
	fmt.Printf("memoSquare(5) = %d\n", memoSquare(5))

	// ========== Bonus: Defer exploration ==========
	fmt.Println("\n=== Bonus: Defer with Arguments ===")
	deferDemo()
}

// deferDemo shows that defer evaluates arguments immediately,
// but executes the function call when the surrounding function returns.
func deferDemo() {
	x := 0
	defer fmt.Printf("deferred with x=%d (captured at defer time)\n", x)
	x = 42
	fmt.Printf("x is now %d\n", x)
	// The deferred call will print x=0, not x=42,
	// because arguments are evaluated when defer is called.
}
