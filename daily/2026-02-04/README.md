# Day 3: Functions, Multiple Returns & Named Returns

Functions are Go's bread and butter. Today you'll master how Go handles them — including some patterns you won't find in most other languages.

## Objectives

1. Write functions with various parameter and return types
2. Use multiple return values (Go's signature move)
3. Understand named returns and when to use them
4. Work with variadic functions
5. Use functions as values (first-class functions)

## Setup

```bash
cd daily/2026-02-04
go mod init funclab
```

## Challenge: Math Toolkit

Build a small math toolkit that shows off Go's function features.

### Part 1: Basic Functions & Multiple Returns

Create `math.go` with these functions:

```go
// Divide returns quotient and remainder.
// Returns an error if divisor is zero.
func Divide(a, b int) (int, int, error)

// MinMax returns the minimum and maximum of a slice.
// Returns an error if the slice is empty.
func MinMax(nums []int) (int, int, error)

// Swap returns its two arguments in reverse order.
func Swap(a, b string) (string, string)
```

Key points to demonstrate:
- Multiple return values (the idiomatic `value, err` pattern)
- Parameter shorthand (`a, b int` instead of `a int, b int`)
- Returning errors instead of panicking

Test in `main.go`:
```
=== Divide ===
17 / 5 = 3 remainder 2
17 / 0 = error: division by zero

=== MinMax ===
[3, 1, 4, 1, 5, 9, 2, 6] → min: 1, max: 9
[] → error: empty slice

=== Swap ===
Before: a=hello, b=world
After:  a=world, b=hello
```

### Part 2: Named Returns

Create `stats.go` with a function that uses **named return values**:

```go
// Stats returns count, sum, and average of a slice of float64.
// Named returns make it clear what each value means.
func Stats(numbers []float64) (count int, sum float64, avg float64, err error)
```

Rules:
- Return an error for empty slices
- Use the **naked return** (just `return`) at least once — then refactor to explicit returns
- In a comment, note which style you prefer and why

Also create a `SummarizeGrades` function:

```go
// SummarizeGrades takes student grades and returns pass/fail counts.
// A passing grade is >= 60.
func SummarizeGrades(grades map[string]int) (passed, failed int)
```

Expected output:
```
=== Stats ===
Numbers: [85.5, 92.3, 78.1, 95.0, 88.7]
Count: 5, Sum: 439.60, Average: 87.92

=== Grades ===
Grades: {Alice: 92, Bob: 55, Charlie: 78, Diana: 43, Eve: 88}
Passed: 3, Failed: 2
```

### Part 3: Variadic Functions

Create `variadic.go`:

```go
// Sum takes any number of ints and returns their total.
func Sum(nums ...int) int

// JoinWith joins strings using a separator.
// First arg is the separator, rest are the strings to join.
func JoinWith(sep string, parts ...string) string

// WrapErrors combines multiple errors into one.
// Skips nil errors. Returns nil if all are nil.
func WrapErrors(errs ...error) error
```

Show how to call variadics both ways:
```go
// Individual arguments
Sum(1, 2, 3, 4, 5)

// Spread a slice
nums := []int{1, 2, 3, 4, 5}
Sum(nums...)
```

Expected output:
```
=== Variadic ===
Sum(1,2,3,4,5) = 15
Sum(spread) = 15
JoinWith(" | ", "Go", "is", "fun") = "Go | is | fun"
WrapErrors(nil, err1, nil, err2) = "multiple errors: file not found; permission denied"
WrapErrors(nil, nil) = <nil>
```

### Part 4: Functions as Values

Create `functional.go`:

```go
// Apply takes a slice of ints and a transform function,
// returns a new slice with the function applied to each element.
func Apply(nums []int, fn func(int) int) []int

// MakeMultiplier returns a function that multiplies by n.
// This is a closure — it captures n.
func MakeMultiplier(n int) func(int) int

// Compose takes two functions and returns their composition.
// Compose(f, g)(x) = f(g(x))
func Compose(f, g func(int) int) func(int) int
```

Expected output:
```
=== Functions as Values ===
Original: [1, 2, 3, 4, 5]
Doubled:  [2, 4, 6, 8, 10]
Squared:  [1, 4, 9, 16, 25]

triple := MakeMultiplier(3)
triple(7) = 21

doubleThenAddOne := Compose(addOne, double)
doubleThenAddOne(5) = 11
```

### Part 5: Putting It Together — Pipeline

Create a small data processing pipeline using what you've built:

```go
// ProcessScores takes raw test scores, filters out invalids (< 0 or > 100),
// applies a curve function, and returns stats.
func ProcessScores(scores []int, curve func(int) int) (count int, sum float64, avg float64, err error)
```

Example usage:
```go
// Curve: add 5 points, cap at 100
curve := func(score int) int {
    s := score + 5
    if s > 100 {
        return 100
    }
    return s
}

scores := []int{72, 85, -1, 91, 150, 68, 77, 95, 83}
// Should filter out -1 and 150, curve the rest
```

Expected output:
```
=== Score Pipeline ===
Raw scores: [72, 85, -1, 91, 150, 68, 77, 95, 83]
Valid scores after curve: [77, 90, 96, 73, 82, 100, 88]
Count: 7, Sum: 606.00, Average: 86.57
```

## Bonus Challenges

- [ ] Create a `Retry` function: `func Retry(attempts int, fn func() error) error` — calls `fn` up to N times until it succeeds
- [ ] Build a `Filter` function: `func Filter(nums []int, predicate func(int) bool) []int`
- [ ] Implement `Reduce`: `func Reduce(nums []int, initial int, fn func(int, int) int) int`
- [ ] Write a `Memoize` wrapper for expensive function calls (hint: use a map in a closure)
- [ ] Explore `defer` with functions — what happens when you defer a function with arguments?

## Key Concepts

- **Multiple returns** — Go functions return multiple values; the `(value, error)` pattern is everywhere
- **Named returns** — document what each return value means; enable naked returns (use sparingly)
- **Variadic functions** — `...T` for flexible argument counts; spread with `slice...`
- **First-class functions** — pass functions around like any other value
- **Closures** — functions that capture variables from their enclosing scope
- **Error returns** — Go uses explicit error returns instead of exceptions

## Common Gotchas

1. Named returns are initialized to zero values — be careful with naked returns in complex functions
2. Variadic parameter must be the last parameter
3. You can't spread a `[]int` into a `...interface{}` — types must match
4. Closures capture variables by reference, not by value (watch out in loops!)
5. Unused function parameters don't cause compile errors, but unused variables still do

## Resources

- [Go Tour: Functions](https://go.dev/tour/basics/4)
- [Go Tour: Multiple Results](https://go.dev/tour/basics/6)
- [Go Tour: Named Returns](https://go.dev/tour/basics/7)
- [Go Tour: Function Closures](https://go.dev/tour/moretypes/25)
- [Effective Go: Functions](https://go.dev/doc/effective_go#functions)
- [Go Blog: Error Handling](https://go.dev/blog/error-handling-and-go)

## Submission

When done:
```bash
git add .
git commit -m "day 3: functions, multiple returns, named returns"
git push
```
