# Day 2: Variables, Types, Constants & Zero Values

**Theme:** Understanding Go's type system from the ground up

## Learning Objectives
- Declare variables using `var`, short declaration (`:=`), and grouped syntax
- Understand Go's static typing and type inference
- Work with basic types: `int`, `float64`, `string`, `bool`, `rune`, `byte`
- Learn constants (`const`) and iota for enumeration
- Master zero values (Go initializes variables to default values)

## Challenge: Temperature Monitor

Build a temperature monitoring CLI that processes readings from different sensors.

### Requirements

1. **Sensor Data Types**
   - Define constants for:
     - `SensorType`: `Celsius`, `Fahrenheit`, `Kelvin` (use `iota`)
     - `Min`, `Max` valid temperature ranges per sensor type
   
2. **Reading Struct**
   ```go
   type Reading struct {
       SensorID    string
       Type        SensorType
       Value       float64
       Timestamp   int64 // Unix timestamp
   }
   ```

3. **Core Functions**
   - `ToCelsius(r Reading) float64`: Convert any reading to Celsius
   - `IsValid(r Reading) bool`: Check if reading is within sensor's valid range
   - `AverageCelsius(readings []Reading) float64`: Average after converting all to Celsius

4. **CLI Program**
   - Initialize a slice with 5+ sample readings (mix of C, F, K)
   - Print each reading: ID, original value, converted value, validity
   - Print average temperature across all readings

### Example Output
```
Sensor A1: 72.5°F → 22.5°C [VALID]
Sensor B2: 300.0K → 26.9°C [VALID]
Sensor C3: 105.0°C → 105.0°C [INVALID: exceeds max]
...
Average: 24.3°C
```

### Rules
- Use `const` with `iota` for sensor types
- Use short declaration (`:=`) where appropriate
- Handle zero values gracefully (empty slice returns 0.0)
- Conversions: `°C = (°F - 32) × 5/9`, `°C = K - 273.15`

## Stretch Goals
- Add a `String()` method to `Reading` for formatted output
- Filter readings by validity before averaging
- Implement min/max temperature detection

## Time Target
30-45 minutes

## References
- [Go by Example: Variables](https://gobyexample.com/variables)
- [Go by Example: Constants](https://gobyexample.com/constants)
- [Effective Go: Names](https://go.dev/doc/effective_go#names)
