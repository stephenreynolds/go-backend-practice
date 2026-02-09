# Day 3: Functions, Multiple Returns & Named Returns

**Theme:** Mastering Go's flexible function signatures

## Learning Objectives
- Write functions with multiple return values
- Use named return values effectively
- Understand variadic functions (`...`)
- Apply the comma-ok idiom for error handling
- Practice function signatures as documentation

## Challenge: File Statistics Calculator

Build a CLI tool that analyzes text files and returns multiple statistics at once.

### Requirements

1. **Core Stats Function**
   ```go
   func AnalyzeText(content string) (lines, words, chars int)
   ```
   - Return multiple stats in a single call
   - Use named returns for clarity
   - Handle empty input gracefully (return 0, 0, 0)

2. **Safe Parse Function**
   ```go
   func ParseInt(s string) (value int, ok bool)
   ```
   - Return a boolean for success/failure (comma-ok pattern)
   - Don't panic on invalid input
   - Used later for parsing command-line arguments

3. **Variadic Sum Function**
   ```go
   func Sum(numbers ...int) int
   ```
   - Accept any number of integers
   - Return total sum

4. **Word Frequency Analysis**
   ```go
   func TopWords(content string, n int) (words []string, counts []int, err error)
   ```
   - Return top N most frequent words
   - Return error if n < 1 or content is empty
   - Multiple return values including error

### CLI Program

```
$ go run main.go sample.txt
=== File Analysis: sample.txt ===
Lines:      42
Words:      318
Characters: 1,847

Top 5 Words:
  1. "the"  → 23 occurrences
  2. "and"  → 15 occurrences
  3. "to"   → 12 occurrences
  ...

Word count sum check: 318 ✓
```

### Test Data
Create a `sample.txt` with any interesting text (a poem, code comments, lorem ipsum — your choice).

### Rules
- Every function must use multiple returns or variadic parameters
- Use named returns where they improve readability
- No external packages — stdlib only (`strings`, `fmt`, `os`, `sort`)
- Handle errors by returning them, not by panicking

## Stretch Goals
- Add `--top N` flag to specify how many top words to show
- Implement `AnalyzeFile(path string) (Stats, error)` that wraps file I/O
- Calculate reading time estimate (avg 200 words/min)
- Support analyzing multiple files with variadic file paths

## Skeleton Code

```go
package main

import (
    "fmt"
    "os"
    "strings"
)

// AnalyzeText returns line, word, and character counts.
// Named returns make the signature self-documenting.
func AnalyzeText(content string) (lines, words, chars int) {
    // Your implementation
    return // naked return uses named values
}

// ParseInt safely parses a string to int, returning ok=false on failure.
func ParseInt(s string) (value int, ok bool) {
    // Your implementation
    return
}

// Sum accepts any number of integers using variadic syntax.
func Sum(numbers ...int) int {
    // Your implementation
    return 0
}

// TopWords returns the n most frequent words and their counts.
func TopWords(content string, n int) (words []string, counts []int, err error) {
    // Your implementation
    return
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <filename>")
        os.Exit(1)
    }
    
    // Read file, analyze, print results
}
```

## Key Patterns to Practice

### Named Returns (self-documenting)
```go
func Divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return // returns 0.0, err
    }
    result = a / b
    return // returns result, nil
}
```

### Comma-OK Idiom
```go
value, ok := ParseInt(input)
if !ok {
    fmt.Println("Invalid number")
}
```

### Variadic Functions
```go
total := Sum(1, 2, 3, 4, 5)
nums := []int{10, 20, 30}
total = Sum(nums...) // spread slice
```

## Time Target
45-60 minutes

## References
- [Go by Example: Functions](https://gobyexample.com/functions)
- [Go by Example: Multiple Return Values](https://gobyexample.com/multiple-return-values)
- [Go by Example: Variadic Functions](https://gobyexample.com/variadic-functions)
- [Effective Go: Functions](https://go.dev/doc/effective_go#functions)
