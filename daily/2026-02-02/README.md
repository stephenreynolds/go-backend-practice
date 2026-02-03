# Day 1: Hello Go

Welcome to Go! Today we set up the environment and write our first programs.

## Objectives

1. Install Go and verify setup
2. Understand Go workspace and modules
3. Write and run your first Go program
4. Learn about packages, imports, and `fmt`

## Setup

```bash
# Check Go is installed
go version
# Should show: go version go1.21+ (or newer)

# Create module for today's challenge
cd daily/2026-02-02
go mod init hello
```

## Challenge

### Part 1: Hello World
Create `main.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Run it:
```bash
go run main.go
```

### Part 2: Explore fmt
Expand your program to demonstrate different `fmt` functions:

1. Use `fmt.Printf` with format verbs (`%s`, `%d`, `%v`)
2. Use `fmt.Sprintf` to format a string without printing
3. Print multiple values with `fmt.Println`

Example output:
```
Hello, Go!
My name is Stephen and I'm learning Go in 2026
Formatted string: Go is awesome!
Multiple: one two three
```

### Part 3: Multiple Files
Create a second file `greet.go` in the same package:
- Define a function `Greet(name string) string` that returns a greeting
- Call it from `main.go`

This demonstrates how Go packages work with multiple files.

## Bonus Challenges

- [ ] Use `fmt.Errorf` to create a formatted error
- [ ] Explore `fmt.Scan` to read user input
- [ ] Check out `go fmt` to auto-format your code
- [ ] Run `go build` to create a binary

## Key Concepts

- **package main**: Entry point package
- **func main()**: Entry point function
- **import**: Brings in other packages
- **fmt**: Format package for I/O
- **:=**: Short variable declaration

## Resources

- [Go Tour: Hello World](https://go.dev/tour/welcome/1)
- [fmt package docs](https://pkg.go.dev/fmt)

## Submission

When done:
```bash
git add .
git commit -m "Add solution to daily 2026-02-02: Hello Go"
git push
```
