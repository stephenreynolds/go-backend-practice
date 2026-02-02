# Day 1: Hello C#

**Estimated Time**: 15-20 minutes  
**Difficulty**: ⭐ Beginner  
**Topic**: Your first C# program

## Objective

Get your C# environment set up and write your first program. No web stuff yet — just the basics.

## Prerequisites

Install the .NET SDK if you haven't:
```bash
# Check if installed
dotnet --version

# If not, install via your package manager or https://dotnet.microsoft.com/download
```

## Challenge

1. **Create a new console project**
   ```bash
   mkdir HelloCSharp && cd HelloCSharp
   dotnet new console
   ```

2. **Edit `Program.cs`** to:
   - Ask the user for their name
   - Greet them by name
   - Tell them the current date and time

3. **Expected output**:
   ```
   What is your name? Stephen
   Hello, Stephen!
   Today is Monday, February 2, 2026 at 1:05 AM.
   ```

## Hints

<details>
<summary>Reading user input</summary>

```csharp
Console.Write("Prompt: ");
string input = Console.ReadLine();
```
</details>

<details>
<summary>Getting current date/time</summary>

```csharp
DateTime now = DateTime.Now;
// or use string formatting
Console.WriteLine($"It is {DateTime.Now}");
```
</details>

<details>
<summary>Formatting dates</summary>

```csharp
DateTime.Now.ToString("dddd, MMMM d, yyyy")  // Monday, February 2, 2026
DateTime.Now.ToString("h:mm tt")              // 1:05 AM
```
</details>

## Bonus Challenges (Optional)

- Add your local timezone to the output
- Ask for their birth year and calculate their age
- Make the greeting change based on time of day ("Good morning", "Good afternoon", etc.)

## Running Your Code

```bash
dotnet run
```

## What You're Learning

- `dotnet new` — Creating projects
- `dotnet run` — Building and running
- `Console.WriteLine()` — Output
- `Console.ReadLine()` — Input
- String interpolation (`$"Hello, {name}!"`)
- `DateTime` — Working with dates/times

## Submission

```bash
git add .
git commit -m "Day 1: Hello C#"
git push
```

Then let me know — I'll take a look!

## Next Up

Tomorrow: Variables, types, and type inference in C#.
