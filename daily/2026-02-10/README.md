# Day 4: Control Flow — if, for, switch

**Theme:** Go's streamlined control structures

## Learning Objectives
- Master `if` with init statements (Go's signature pattern)
- Use all three forms of `for` loops (C-style, range, infinite)
- Write clean `switch` statements with multiple cases
- Understand when `break`, `continue`, and `fallthrough` apply
- Avoid common control flow anti-patterns

## Challenge: Log File Analyzer

Build a CLI tool that parses a server log file and generates statistics using only control flow constructs.

### The Log Format

Each line follows this pattern:
```
[LEVEL] TIMESTAMP MESSAGE
```

Example `server.log`:
```
[INFO] 2026-02-10T09:00:01 Server started on port 8080
[DEBUG] 2026-02-10T09:00:02 Loading configuration from config.yaml
[INFO] 2026-02-10T09:01:15 User alice logged in
[WARN] 2026-02-10T09:02:33 High memory usage: 85%
[ERROR] 2026-02-10T09:03:44 Database connection timeout after 30s
[INFO] 2026-02-10T09:04:01 User bob logged in
[ERROR] 2026-02-10T09:05:22 Failed to process request: invalid JSON
[DEBUG] 2026-02-10T09:06:00 Cache hit ratio: 0.73
[WARN] 2026-02-10T09:07:18 Rate limit approaching for IP 192.168.1.50
[INFO] 2026-02-10T09:08:45 User alice logged out
[FATAL] 2026-02-10T09:10:00 Out of memory - shutting down
```

### Requirements

1. **Parse and Count by Level**
   Use `switch` to categorize each log level:
   ```go
   switch level {
   case "INFO":
       // ...
   case "WARN", "WARNING":
       // handle both variants
   case "ERROR", "FATAL":
       // group critical issues
   default:
       // unknown levels
   }
   ```

2. **Filter by Severity**
   Implement `--min-level` flag:
   - `DEBUG` → show all
   - `INFO` → skip DEBUG
   - `WARN` → skip DEBUG, INFO
   - `ERROR` → only ERROR and FATAL
   
   Use if-init pattern to parse the flag:
   ```go
   if level := getMinLevel(); level != "" {
       // filter logic
   }
   ```

3. **Time Range Filter**
   Parse timestamps and filter logs within a time window:
   ```go
   for i := 0; i < len(logs); i++ {
       // C-style when you need the index
   }
   
   for _, log := range logs {
       // range when you just need values
   }
   
   for {
       // infinite loop with break condition
   }
   ```

4. **Pattern Search**
   Search for keywords in messages, printing line numbers:
   ```go
   for lineNum, line := range lines {
       if strings.Contains(line, keyword) {
           // found match
       }
   }
   ```

### CLI Usage

```
$ go run main.go server.log
=== Log Analysis ===

Level Counts:
  DEBUG:  2
  INFO:   4
  WARN:   2
  ERROR:  2
  FATAL:  1
  
Total: 11 entries
Critical (ERROR+FATAL): 3

$ go run main.go server.log --min-level WARN
[WARN] 2026-02-10T09:02:33 High memory usage: 85%
[ERROR] 2026-02-10T09:03:44 Database connection timeout after 30s
[ERROR] 2026-02-10T09:05:22 Failed to process request: invalid JSON
[WARN] 2026-02-10T09:07:18 Rate limit approaching for IP 192.168.1.50
[FATAL] 2026-02-10T09:10:00 Out of memory - shutting down

Showing 5 entries (WARN and above)

$ go run main.go server.log --search "alice"
Line 3: [INFO] 2026-02-10T09:01:15 User alice logged in
Line 10: [INFO] 2026-02-10T09:08:45 User alice logged out

Found 2 matches for "alice"
```

### Rules
- Use all three `for` variants at least once
- Use `switch` (not if-else chains) for level handling
- Use at least one if-init statement
- No `goto` (ever)
- stdlib only (`fmt`, `os`, `strings`, `bufio`)

## Skeleton Code

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Log levels in order of severity
var severityOrder = map[string]int{
	"DEBUG": 0,
	"INFO":  1,
	"WARN":  2,
	"ERROR": 3,
	"FATAL": 4,
}

// parseLine extracts level and message from a log line
// Returns empty strings if format is invalid
func parseLine(line string) (level, timestamp, message string) {
	// Hint: look for [ and ] to extract level
	// Your implementation
	return
}

// shouldShow returns true if the log level meets minimum severity
func shouldShow(level, minLevel string) bool {
	// Use if-init with map lookup
	// Your implementation
	return true
}

// countByLevel tallies occurrences of each log level
func countByLevel(lines []string) map[string]int {
	counts := make(map[string]int)
	// Use for-range
	// Use switch for level categorization
	// Your implementation
	return counts
}

// filterByLevel returns only lines meeting minimum severity
func filterByLevel(lines []string, minLevel string) []string {
	var result []string
	// Use for-range and continue to skip
	// Your implementation
	return result
}

// searchLines finds lines containing keyword
func searchLines(lines []string, keyword string) []struct {
	lineNum int
	text    string
} {
	var matches []struct {
		lineNum int
		text    string
	}
	// Use for with index (C-style or range with index)
	// Your implementation
	return matches
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <logfile> [--min-level LEVEL] [--search KEYWORD]")
		os.Exit(1)
	}

	filename := os.Args[1]
	
	// Parse flags manually (good control flow practice)
	var minLevel, searchKeyword string
	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--min-level":
			if i+1 < len(os.Args) {
				minLevel = os.Args[i+1]
				i++ // skip next arg
			}
		case "--search":
			if i+1 < len(os.Args) {
				searchKeyword = os.Args[i+1]
				i++
			}
		}
	}

	// Read file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Main logic based on mode
	// Your implementation
}
```

## Key Patterns to Practice

### If with Init Statement
```go
// Scope 'err' to just this block
if err := doSomething(); err != nil {
    return err
}

// Scope map lookup to the condition
if val, ok := myMap[key]; ok {
    fmt.Println(val)
}
```

### For Loop Variants
```go
// C-style (when you need precise index control)
for i := 0; i < n; i++ { }

// Range over slice (idiomatic for iteration)
for index, value := range slice { }

// Range over map
for key, value := range myMap { }

// Infinite with break
for {
    if done {
        break
    }
}

// While-style
for condition {
    // body
}
```

### Switch Best Practices
```go
// Multiple values per case
switch day {
case "Sat", "Sun":
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}

// No fallthrough by default (unlike C)
// Use explicit 'fallthrough' if needed (rare)

// Switch with no expression (cleaner than if-else chains)
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
default:
    grade = "C"
}
```

## Stretch Goals
- Add `--after` and `--before` flags to filter by timestamp
- Implement `--tail N` to show last N entries (requires different loop strategy)
- Add colorized output (ERROR in red, WARN in yellow)
- Count unique users mentioned in login/logout messages
- Calculate time between first and last log entry

## Test Data

Create `server.log` with the example content above, or generate your own with varied levels and timestamps.

## Time Target
45-60 minutes

## References
- [Go by Example: If/Else](https://gobyexample.com/if-else)
- [Go by Example: For](https://gobyexample.com/for)
- [Go by Example: Switch](https://gobyexample.com/switch)
- [Effective Go: Control Structures](https://go.dev/doc/effective_go#control-structures)
