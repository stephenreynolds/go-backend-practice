# Day 2: Variables, Types, Constants & Zero Values

Today you'll get comfortable with Go's type system — how variables work, what types exist, how constants differ from variables, and the important concept of zero values.

## Objectives

1. Declare variables using `var`, `:=`, and grouped declarations
2. Work with Go's basic types (int, float64, string, bool, byte, rune)
3. Understand type conversions (Go has no implicit casting!)
4. Define and use constants, including `iota`
5. Understand zero values and why they matter

## Setup

```bash
cd daily/2026-02-03
go mod init types-lab
```

## Challenge: Unit Converter

Build a small unit converter that demonstrates Go's type system in action.

### Part 1: Temperature Converter

Create `main.go` with functions that convert between temperature scales:

```go
func celsiusToFahrenheit(c float64) float64
func fahrenheitToCelsius(f float64) float64
func celsiusToKelvin(c float64) float64
```

Requirements:
- Use **named constants** for the conversion factors (e.g., `const absoluteZero = -273.15`)
- Print results using `fmt.Printf` with `%.2f` formatting
- Demonstrate the difference between `var` declaration and `:=` short declaration

Expected output:
```
=== Temperature Converter ===
100.00°C = 212.00°F
72.00°F = 22.22°C
0.00°C = 273.15K
Absolute zero: -273.15°C = 0.00K
```

### Part 2: Type-Safe Distance Types

Create `distance.go` with **custom types** based on float64:

```go
type Meters float64
type Feet float64
type Miles float64
```

Add conversion functions:
```go
func (m Meters) ToFeet() Feet
func (m Meters) ToMiles() Miles
func (f Feet) ToMeters() Meters
```

This demonstrates:
- Custom types based on built-in types
- Methods on types
- Type safety (you can't accidentally pass `Meters` where `Feet` is expected)

Try this in `main()` and see what happens:
```go
var m Meters = 100
var f Feet = 328.084
// This should NOT compile:
// f = m  // type mismatch!
// But this works:
f = m.ToFeet()
```

### Part 3: Measurement Enum with iota

Create `units.go` with a constant enum for measurement systems:

```go
type MeasurementSystem int

const (
    Metric MeasurementSystem = iota
    Imperial
    Nautical
)
```

Add a `String()` method so it prints nicely:
```go
func (ms MeasurementSystem) String() string
```

Create a function that returns the appropriate unit labels:
```go
func UnitLabels(sys MeasurementSystem) (distance, weight, volume string)
```

This demonstrates:
- `iota` for enumerations
- Multiple return values (sneak peek from tomorrow)
- The `Stringer` interface

Expected output:
```
=== Measurement Systems ===
System: Metric → distance: meter, weight: kilogram, volume: liter
System: Imperial → distance: foot, weight: pound, volume: gallon
System: Nautical → distance: nautical mile, weight: kilogram, volume: liter
```

### Part 4: Zero Value Explorer

Create `zeros.go` with a function that prints the zero value of every basic type:

```go
func showZeroValues() {
    var i int
    var f float64
    var b bool
    var s string
    var r rune
    var by byte
    // ... print them all with their types using %v and %T
}
```

Expected output:
```
=== Zero Values ===
int:     0       (type: int)
float64: 0       (type: float64)
bool:    false   (type: bool)
string:  ""      (type: string)
rune:    0       (type: int32)
byte:    0       (type: uint8)
```

**Key insight:** In Go, every variable has a useful zero value. You never have `nil` for basic types. This is a deliberate design choice — understand why it matters.

## Bonus Challenges

- [ ] Add a `Weight` type with Kilograms/Pounds and conversions
- [ ] Use `const` block with `iota` to define byte size units (KB, MB, GB, TB) where each is 1024× the previous (hint: `1 << (10 * iota)`)
- [ ] Explore what happens with integer overflow: `var x uint8 = 255; x++`
- [ ] Try type conversion between `int` and `float64` — what gets lost?
- [ ] Use `math.MaxInt64` and `math.MinInt64` to explore type bounds

## Key Concepts

- **var vs :=** — `var` works everywhere, `:=` only inside functions
- **Zero values** — Go always initializes: 0, false, "", nil
- **No implicit conversions** — you must explicitly convert between types
- **Custom types** — create new types from existing ones for type safety
- **iota** — auto-incrementing constant generator
- **rune vs byte** — rune is `int32` (Unicode), byte is `uint8`

## Common Gotchas

1. `float64` arithmetic isn't exact: `0.1 + 0.2 != 0.3`
2. Integer division truncates: `7 / 2 == 3` (not 3.5)
3. You can't add `int` and `float64` without conversion
4. Constants in Go are untyped until used — they have arbitrary precision

## Resources

- [Go Tour: Types](https://go.dev/tour/basics/11)
- [Go Spec: Constants](https://go.dev/ref/spec#Constants)
- [Go Blog: Constants](https://go.dev/blog/constants)
- [Effective Go: Variables](https://go.dev/doc/effective_go#variables)

## Submission

When done:
```bash
git add .
git commit -m "day 2: variables, types, constants & zero values"
git push
```
