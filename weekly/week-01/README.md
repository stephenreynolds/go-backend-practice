# Week 1 Project: CLI Task Manager

Build a command-line task manager to practice Go fundamentals.

## Requirements

A CLI tool called `tasks` that supports:

```bash
# Add a task
tasks add "Buy groceries"
tasks add "Learn Go"

# List all tasks
tasks list
# Output:
# 1. [ ] Buy groceries
# 2. [ ] Learn Go

# Complete a task
tasks done 1

# List again
tasks list
# Output:
# 1. [x] Buy groceries
# 2. [ ] Learn Go

# Delete a task
tasks delete 1
```

## Specifications

- Tasks stored in a JSON file (`~/.tasks.json`)
- Each task has: id, title, completed, created_at
- Handle errors gracefully (file not found, invalid ID, etc.)
- Use `os.Args` for command-line arguments

## Stretch Goals

- [ ] Add due dates to tasks
- [ ] Filter by completed/incomplete
- [ ] Add priority levels
- [ ] Color output (completed = green, overdue = red)
- [ ] Use a CLI framework like `cobra` or `urfave/cli`

## Suggested Structure

```
weekly/week-01/
├── main.go           # Entry point, argument parsing
├── task.go           # Task struct and methods
├── storage.go        # File read/write operations
└── go.mod
```

## Start Date
Begin after completing Day 5 (or earlier if you want!)

## Resources

- [os.Args](https://pkg.go.dev/os#Args)
- [encoding/json](https://pkg.go.dev/encoding/json)
- [os package](https://pkg.go.dev/os)
