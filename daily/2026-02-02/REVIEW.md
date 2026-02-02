# Day 1 Review: Hello C#

**Submitted**: 2026-02-02  
**Score**: 100/100 ✓

## Solution

```csharp
Console.Write("What is your name? ");
string? name = Console.ReadLine();

if (name is null)
{
    Console.WriteLine("Nothing was read. Exiting.");
    return;
}

Console.WriteLine($"Hello, {name}!");

var datetime = DateTime.Now;
Console.WriteLine(
    $"Today is {datetime.ToString("dddd, MMMM d, yyyy")} at {datetime.ToString("h:mm tt")}."
);
```

## What You Did Well

- **Nullable `string?`** — Shows awareness of null safety in modern C#
- **`is null` pattern matching** — Clean, idiomatic C# style
- **Early return** — Good defensive programming practice
- **String interpolation** — Correct usage with `$"...{variable}..."`
- **Date formatting** — Exactly right format strings

## Suggestions

The `:format` syntax inside interpolation is a nice shorthand — no need for `.ToString()`:

```csharp
// Option 1: Inline formatting
Console.WriteLine($"Today is {DateTime.Now:dddd, MMMM d, yyyy} at {DateTime.Now:h:mm tt}.");

// Option 2: Capture once, format inline
var now = DateTime.Now;
Console.WriteLine($"Today is {now:dddd, MMMM d, yyyy} at {now:h:mm tt}.");
```

This is purely stylistic — your solution is correct and readable.

## Concepts Demonstrated

- ✅ Console I/O (`Write`, `WriteLine`, `ReadLine`)
- ✅ Nullable reference types (`string?`)
- ✅ Pattern matching (`is null`)
- ✅ String interpolation
- ✅ DateTime formatting
- ✅ Early return pattern

## Next Up

Day 2: Variables & Types — `int`, `string`, `bool`, `var`, type inference
