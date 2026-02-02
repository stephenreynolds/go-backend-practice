# Day 1: Minimal API Basics

**Estimated Time**: 20-30 minutes  
**Difficulty**: ⭐ Beginner  
**Topic**: ASP.NET Core Minimal APIs

## Objective

Create a minimal API that acts as a simple calculator. No controllers, no MVC — just clean, minimal endpoints.

## Requirements

1. Create a new minimal API project
2. Implement these endpoints:

| Method | Route | Description | Example |
|--------|-------|-------------|---------|
| GET | `/add/{a}/{b}` | Add two numbers | `/add/5/3` → `8` |
| GET | `/subtract/{a}/{b}` | Subtract b from a | `/subtract/10/4` → `6` |
| GET | `/multiply/{a}/{b}` | Multiply two numbers | `/multiply/6/7` → `42` |
| GET | `/divide/{a}/{b}` | Divide a by b | `/divide/20/4` → `5` |

3. Handle edge cases:
   - Division by zero should return 400 Bad Request with a message
   - Non-numeric inputs should return 400 Bad Request

4. Add a root endpoint `GET /` that returns API info:
   ```json
   {
     "name": "Calculator API",
     "version": "1.0",
     "endpoints": ["/add", "/subtract", "/multiply", "/divide"]
   }
   ```

## Bonus Challenges (Optional)

- Add a `POST /calculate` endpoint that accepts JSON:
  ```json
  { "operation": "add", "a": 5, "b": 3 }
  ```
- Add request logging middleware
- Return results as JSON objects: `{ "result": 8, "operation": "add" }`

## Getting Started

```bash
dotnet new webapi -minimal -n Calculator
cd Calculator
# Edit Program.cs
dotnet run
```

## Submission

When complete:
```bash
git add .
git commit -m "Day 1: Minimal API Calculator"
git push
```

Then let me know and I'll review your code!

## Learning Resources

- [Minimal APIs Overview](https://learn.microsoft.com/en-us/aspnet/core/fundamentals/minimal-apis/overview)
- [Route Parameters](https://learn.microsoft.com/en-us/aspnet/core/fundamentals/minimal-apis/parameter-binding)
- [Return Types](https://learn.microsoft.com/en-us/aspnet/core/fundamentals/minimal-apis/responses)
