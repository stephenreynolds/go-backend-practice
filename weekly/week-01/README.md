# Week 1 Project: TODO API

**Estimated Time**: 3-5 hours total  
**Deadline**: End of Week 1  
**Difficulty**: ⭐⭐ Beginner-Intermediate

## Objective

Build a complete TODO API using ASP.NET Core Minimal APIs with in-memory storage. This project consolidates all the daily lessons from Week 1.

## Requirements

### Core Features

1. **CRUD Operations**
   - `GET /todos` — List all todos (with optional filtering)
   - `GET /todos/{id}` — Get a specific todo
   - `POST /todos` — Create a new todo
   - `PUT /todos/{id}` — Update a todo
   - `DELETE /todos/{id}` — Delete a todo

2. **Todo Model**
   ```csharp
   public record Todo
   {
       public int Id { get; init; }
       public string Title { get; init; }
       public string? Description { get; init; }
       public bool IsComplete { get; init; }
       public DateTime CreatedAt { get; init; }
       public DateTime? CompletedAt { get; init; }
   }
   ```

3. **Filtering & Sorting** (GET /todos)
   - `?completed=true` — Filter by completion status
   - `?search=keyword` — Search in title/description
   - `?sortBy=createdAt&order=desc` — Sorting

4. **Validation**
   - Title is required, 1-200 characters
   - Return 400 with validation errors

5. **Error Handling**
   - 404 for missing todos
   - 400 for validation errors
   - Use ProblemDetails format

### Architecture Requirements

- Use Dependency Injection for the todo "repository"
- Create a `ITodoService` interface
- Implement an `InMemoryTodoService`
- Use DTOs for input/output (don't expose internal models directly)
- Add request/response logging middleware

### Endpoints Summary

| Method | Route | Description | Response |
|--------|-------|-------------|----------|
| GET | /todos | List todos | 200 + array |
| GET | /todos/{id} | Get todo | 200 or 404 |
| POST | /todos | Create todo | 201 + location |
| PUT | /todos/{id} | Update todo | 200 or 404 |
| DELETE | /todos/{id} | Delete todo | 204 or 404 |
| PATCH | /todos/{id}/complete | Mark complete | 200 or 404 |

## Bonus Challenges

- Add pagination (`?page=1&pageSize=10`)
- Add bulk operations (`DELETE /todos` with body)
- Add OpenAPI documentation with Swagger
- Add request rate limiting
- Add response caching

## Project Structure

```
TodoApi/
├── Program.cs           # App entry, DI setup, endpoints
├── Models/
│   └── Todo.cs
├── DTOs/
│   ├── CreateTodoDto.cs
│   ├── UpdateTodoDto.cs
│   └── TodoResponseDto.cs
├── Services/
│   ├── ITodoService.cs
│   └── InMemoryTodoService.cs
└── Middleware/
    └── RequestLoggingMiddleware.cs
```

## Getting Started

```bash
dotnet new webapi -minimal -n TodoApi
cd TodoApi
# Start building...
dotnet run
```

## Submission

When complete:
```bash
git add .
git commit -m "Week 1: TODO API with in-memory storage"
git push
```

## Grading Criteria

| Criteria | Points |
|----------|--------|
| All CRUD endpoints work | 25 |
| Proper HTTP status codes | 15 |
| Validation implemented | 15 |
| DI and service pattern | 15 |
| DTOs used correctly | 10 |
| Error handling | 10 |
| Code organization | 10 |
| **Bonus**: Extra features | +5 each |

**Total**: 100 points (+bonus)
