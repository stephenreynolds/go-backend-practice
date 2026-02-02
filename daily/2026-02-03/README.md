# Day 2: Variables & Types

**Phase 0, Week 1** — C# Basics  
**Date**: February 3, 2026  
**Estimated Time**: 30-45 minutes

---

## Learning Goals

- Declare variables with explicit types and `var`
- Understand C# primitive types: `int`, `string`, `bool`, `double`, `char`
- Use type inference appropriately
- Convert between types safely
- Work with `const` and `readonly`

---

## Quick Reference

```csharp
// Explicit typing
int age = 30;
string name = "Stephen";
bool isActive = true;
double price = 19.99;
char grade = 'A';

// Type inference with var
var count = 42;          // inferred as int
var message = "Hello";   // inferred as string

// Constants (compile-time)
const double Pi = 3.14159;

// Type conversion
int num = 42;
string text = num.ToString();
int parsed = int.Parse("123");
bool success = int.TryParse("abc", out int result);  // safe parsing
```

---

## Challenge: Temperature Converter

Build a console app that converts temperatures between Celsius and Fahrenheit.

### Requirements

1. **Create a new console project** called `TemperatureConverter`
   ```bash
   dotnet new console -n TemperatureConverter
   cd TemperatureConverter
   ```

2. **Implement the following features**:
   - Prompt the user for a temperature value
   - Prompt for the unit (C or F)
   - Convert to the other unit and display the result
   - Handle invalid input gracefully (don't crash!)

3. **Formulas**:
   - Celsius to Fahrenheit: `F = (C × 9/5) + 32`
   - Fahrenheit to Celsius: `C = (F - 32) × 5/9`

4. **Output format**:
   ```
   Enter temperature: 100
   Enter unit (C/F): C
   100°C = 212°F
   ```

### Stretch Goals (optional)

- [ ] Support lowercase input (`c` and `f`)
- [ ] Add Kelvin conversion (`K = C + 273.15`)
- [ ] Loop until user types "quit"
- [ ] Show rounded results (2 decimal places)

---

## Starter Code

```csharp
// Program.cs
Console.WriteLine("=== Temperature Converter ===");

// TODO: Prompt for temperature value

// TODO: Prompt for unit (C or F)

// TODO: Convert and display result
```

---

## Hints

<details>
<summary>Hint 1: Reading user input</summary>

```csharp
Console.Write("Enter temperature: ");
string input = Console.ReadLine() ?? "";
```

</details>

<details>
<summary>Hint 2: Safe parsing</summary>

```csharp
if (double.TryParse(input, out double temp))
{
    // temp now holds the parsed value
}
else
{
    Console.WriteLine("Invalid number!");
}
```

</details>

<details>
<summary>Hint 3: String comparison (case-insensitive)</summary>

```csharp
string unit = Console.ReadLine()?.ToUpper() ?? "";
if (unit == "C")
{
    // Celsius input
}
```

</details>

---

## Submission

1. Complete the challenge in `TemperatureConverter/Program.cs`
2. Test with various inputs:
   - `100 C` → should output `212°F`
   - `32 F` → should output `0°C`
   - Invalid input → should show error message
3. Commit your work:
   ```bash
   git add .
   git commit -m "Day 2: Temperature Converter"
   ```

---

## Concepts to Review

After completing this challenge, make sure you understand:

- [ ] When to use `var` vs explicit types
- [ ] The difference between `Parse()` and `TryParse()`
- [ ] Why `double` is better than `int` for calculations with decimals
- [ ] How `const` differs from regular variables

---

## Tomorrow's Preview

**Day 3: Control Flow** — We'll add branching logic with `if`, `switch`, and loops. Your temperature converter will get an upgrade!
