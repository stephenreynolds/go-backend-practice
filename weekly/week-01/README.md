# Week 1 Project: Task List CLI

**Estimated Time**: 2-3 hours total  
**Difficulty**: ⭐ Beginner  
**Topic**: C# Fundamentals in Practice

## Objective

Build a simple command-line task manager. No databases, no web — just a console app that manages a list of tasks in memory.

## Requirements

### Core Features

1. **Menu-driven interface**
   ```
   === Task Manager ===
   1. Add task
   2. List tasks
   3. Complete task
   4. Remove task
   5. Exit
   
   Choose an option: 
   ```

2. **Add tasks** — Prompt for task description, add to list

3. **List tasks** — Show all tasks with numbers and completion status
   ```
   Your Tasks:
   1. [ ] Buy groceries
   2. [x] Finish homework
   3. [ ] Call mom
   ```

4. **Complete task** — Mark a task as done by number

5. **Remove task** — Delete a task by number

6. **Exit** — Quit the program

### Task Data

Each task should have:
- Description (text)
- Completed status (yes/no)

### Edge Cases to Handle

- Empty task list ("No tasks yet!")
- Invalid menu choice
- Invalid task number
- Empty description

## Bonus Challenges

- Save tasks to a file, load on startup
- Add due dates to tasks
- Show only incomplete tasks
- Search tasks by keyword
- Add priority levels (high, medium, low)

## Getting Started

```bash
cd weekly/week-01
dotnet new console -n TaskList
cd TaskList
# Start coding in Program.cs
```

## Suggested Approach

1. Start with a simple `List<string>` for task descriptions
2. Build the menu loop first
3. Add each feature one at a time
4. Refactor to use a `Task` class if you want (optional for week 1)

## Example Session

```
=== Task Manager ===
1. Add task
2. List tasks
3. Complete task
4. Remove task
5. Exit

Choose: 1
Enter task: Buy milk
Added: Buy milk

Choose: 1  
Enter task: Read C# book
Added: Read C# book

Choose: 2
Your Tasks:
1. [ ] Buy milk
2. [ ] Read C# book

Choose: 3
Task number to complete: 1
Completed: Buy milk

Choose: 2
Your Tasks:
1. [x] Buy milk
2. [ ] Read C# book

Choose: 5
Goodbye!
```

## What You're Learning

- Console I/O (`ReadLine`, `WriteLine`)
- Control flow (`while`, `switch` or `if/else`)
- Collections (`List<T>`)
- Basic program structure
- Input validation

## Submission

```bash
git add .
git commit -m "Week 1: Task List CLI"
git push
```

## Grading

| Criteria | Points |
|----------|--------|
| Menu loop works | 20 |
| Add task works | 20 |
| List tasks works | 20 |
| Complete task works | 15 |
| Remove task works | 15 |
| Handles edge cases | 10 |
| **Bonus features** | +10 each |

**Total**: 100 points (+bonus)
