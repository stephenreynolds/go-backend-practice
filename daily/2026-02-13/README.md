# Day 7: Pointers and Memory

**Theme:** Understanding Go's memory model — when to share, when to copy

## Learning Objectives
- Master the `&` (address-of) and `*` (dereference) operators
- Understand value semantics vs pointer semantics
- Know when to use pointer receivers vs value receivers
- Handle nil pointers safely
- Recognize common pointer pitfalls (aliasing, loop variables)

## Challenge: Database Connection Pool

Build a connection pool manager — a pattern used in every production backend. This naturally exercises pointers because connections are shared resources that must be passed by reference, tracked, and properly released.

### Core Concepts

```go
// The & operator: get the address of a value
x := 42
ptr := &x        // ptr is *int, holds address of x
fmt.Println(ptr) // 0xc000018090 (memory address)

// The * operator: dereference (access value at address)
fmt.Println(*ptr) // 42
*ptr = 100        // modify the value at that address
fmt.Println(x)    // 100 — x changed!

// Pointers as function parameters (shared mutation)
func double(n *int) {
    *n = *n * 2
}
val := 5
double(&val)
fmt.Println(val) // 10

// Value receivers (copy) vs Pointer receivers (shared)
type Counter struct { n int }

func (c Counter) ValueIncrement() { c.n++ }    // modifies copy, useless!
func (c *Counter) PointerIncrement() { c.n++ } // modifies original

c := Counter{n: 0}
c.ValueIncrement()   // c.n still 0
c.PointerIncrement() // c.n now 1

// nil pointer check
var ptr *int = nil
if ptr != nil {
    fmt.Println(*ptr)  // safe
}
// *ptr without check = panic!

// new() allocates zeroed memory and returns pointer
ptr := new(int)      // *int pointing to 0
*ptr = 42
```

### Requirements

1. **Connection Structure**
   ```go
   type Connection struct {
       ID        int
       Host      string
       Port      int
       InUse     bool
       CreatedAt time.Time
       LastUsed  time.Time
       QueryCount int
   }
   
   // Methods (think: value or pointer receiver?)
   func (c *Connection) Execute(query string) (string, error)
   func (c *Connection) IsHealthy() bool
   func (c *Connection) MarkUsed()
   func (c Connection) String() string  // value is fine here — why?
   ```

2. **Pool Structure**
   ```go
   type Pool struct {
       connections []*Connection  // slice of pointers — why not []Connection?
       maxSize     int
       host        string
       port        int
       nextID      int
   }
   
   type PoolStats struct {
       Total     int
       InUse     int
       Available int
       MaxSize   int
   }
   ```

3. **Pool Operations**
   - `NewPool(host string, port int, maxSize int) *Pool` — create pool
   - `Acquire() (*Connection, error)` — get a connection (reuse or create)
   - `Release(conn *Connection)` — return connection to pool
   - `Close()` — close all connections
   - `Stats() PoolStats` — pool statistics
   - `HealthCheck() int` — close unhealthy connections, return count closed

4. **Connection Lifecycle**
   - Pool starts empty, creates connections on demand up to `maxSize`
   - `Acquire()` first tries to find an available connection
   - If none available and pool isn't full, create new connection
   - If pool is full and all in use, return error
   - `Release()` marks connection as available for reuse
   - Connections track `QueryCount` and `LastUsed`

5. **Health Checking**
   - Connection is "unhealthy" if `LastUsed` was more than 5 minutes ago
   - `HealthCheck()` removes unhealthy connections from pool

### CLI Interface

```
$ go run main.go
Connection Pool Demo

Commands:
  acquire              - Get a connection from pool
  release <id>         - Return connection to pool
  execute <id> <query> - Run query on connection
  list                 - Show all connections
  stats                - Show pool statistics
  health               - Run health check
  close                - Close all connections
  quit                 - Exit

> acquire
Acquired connection #1 (localhost:5432)

> acquire
Acquired connection #2 (localhost:5432)

> list
Connections:
  [1] localhost:5432 — IN USE (queries: 0)
  [2] localhost:5432 — IN USE (queries: 0)

> execute 1 SELECT * FROM users
Connection #1: Executed "SELECT * FROM users"
Result: "Query executed successfully" (simulated)

> execute 1 SELECT * FROM orders
Connection #1: Executed "SELECT * FROM orders"
Result: "Query executed successfully" (simulated)

> release 1
Released connection #1

> list
Connections:
  [1] localhost:5432 — available (queries: 2)
  [2] localhost:5432 — IN USE (queries: 0)

> stats
Pool Statistics:
  Total:     2
  In Use:    1
  Available: 1
  Max Size:  5

> acquire
Reusing connection #1 (localhost:5432)

> list
Connections:
  [1] localhost:5432 — IN USE (queries: 2)
  [2] localhost:5432 — IN USE (queries: 0)

> health
Health check: 0 connections removed

> close
Closed all connections
```

### Key Patterns to Master

**Why `[]*Connection` instead of `[]Connection`?**
```go
// With []Connection — copying creates independent objects
conns := []Connection{{ID: 1, InUse: false}}
c := conns[0]      // c is a COPY
c.InUse = true     // modifies the copy
fmt.Println(conns[0].InUse) // false! Original unchanged

// With []*Connection — all references point to same object
conns := []*Connection{{ID: 1, InUse: false}}
c := conns[0]      // c points to same Connection
c.InUse = true     // modifies the shared object
fmt.Println(conns[0].InUse) // true! Works as expected
```

**Pointer receivers for mutation:**
```go
func (c *Connection) MarkUsed() {
    c.LastUsed = time.Now()
    c.QueryCount++
}

// Called on a pointer (automatic dereference)
conn := &Connection{ID: 1}
conn.MarkUsed()  // Go converts to (*conn).MarkUsed() automatically
```

**Searching and returning pointers:**
```go
func (p *Pool) findAvailable() *Connection {
    for _, conn := range p.connections {
        if !conn.InUse {
            return conn  // return the pointer, not a copy
        }
    }
    return nil  // no available connection
}

func (p *Pool) Acquire() (*Connection, error) {
    if conn := p.findAvailable(); conn != nil {
        conn.InUse = true  // modify shared state
        conn.MarkUsed()
        return conn, nil
    }
    // create new or return error...
}
```

**Nil pointer safety:**
```go
func (p *Pool) Release(conn *Connection) {
    if conn == nil {
        return  // guard against nil
    }
    conn.InUse = false
    conn.LastUsed = time.Now()
}
```

**Common pitfall — loop variable capture:**
```go
// WRONG: all pointers will point to last element!
var ptrs []*Connection
for _, c := range connections {
    ptrs = append(ptrs, &c)  // &c is same address each iteration!
}

// RIGHT: capture the index or create explicit copy
for i := range connections {
    ptrs = append(ptrs, &connections[i])
}
// OR
for _, c := range connections {
    c := c  // shadow with new variable
    ptrs = append(ptrs, &c)
}
```

### Rules
- Connections must be passed by pointer to maintain shared state
- All mutations must use pointer receivers
- Handle nil pointers gracefully (never panic)
- Use value receivers only for read-only methods
- stdlib only (`fmt`, `bufio`, `os`, `strings`, `time`, `errors`)

## Skeleton Code

```go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	ErrPoolExhausted = errors.New("connection pool exhausted")
	ErrConnNotFound  = errors.New("connection not found")
)

type Connection struct {
	ID         int
	Host       string
	Port       int
	InUse      bool
	CreatedAt  time.Time
	LastUsed   time.Time
	QueryCount int
}

// Execute runs a query on this connection (simulated)
func (c *Connection) Execute(query string) (string, error) {
	// Your implementation
	// - Check if connection is in use (return error if not)
	// - Update LastUsed and QueryCount
	// - Return simulated result
	return "", nil
}

// IsHealthy returns true if connection was used recently (within 5 min)
func (c *Connection) IsHealthy() bool {
	// Your implementation
	return false
}

// MarkUsed updates LastUsed timestamp
func (c *Connection) MarkUsed() {
	// Your implementation
}

// String returns a human-readable representation
// Question: Why is a value receiver okay here?
func (c Connection) String() string {
	status := "available"
	if c.InUse {
		status = "IN USE"
	}
	return fmt.Sprintf("[%d] %s:%d — %s (queries: %d)",
		c.ID, c.Host, c.Port, status, c.QueryCount)
}

type PoolStats struct {
	Total     int
	InUse     int
	Available int
	MaxSize   int
}

type Pool struct {
	connections []*Connection
	maxSize     int
	host        string
	port        int
	nextID      int
}

// NewPool creates a new connection pool
func NewPool(host string, port int, maxSize int) *Pool {
	return &Pool{
		connections: make([]*Connection, 0),
		maxSize:     maxSize,
		host:        host,
		port:        port,
		nextID:      1,
	}
}

// createConnection creates a new connection (internal use)
func (p *Pool) createConnection() *Connection {
	// Your implementation
	// - Create connection with next ID
	// - Increment nextID
	// - Set host, port, timestamps
	// - Append to pool
	return nil
}

// findAvailable finds an available (not in use) connection
func (p *Pool) findAvailable() *Connection {
	// Your implementation
	return nil
}

// findByID finds a connection by its ID
func (p *Pool) findByID(id int) *Connection {
	// Your implementation
	return nil
}

// Acquire gets a connection from the pool
func (p *Pool) Acquire() (*Connection, error) {
	// Your implementation
	// 1. Try to find available connection
	// 2. If none, create new if under maxSize
	// 3. If at maxSize, return ErrPoolExhausted
	// Remember to mark as InUse!
	return nil, nil
}

// Release returns a connection to the pool
func (p *Pool) Release(conn *Connection) {
	// Your implementation
	// - Handle nil gracefully
	// - Mark as not in use
	// - Update LastUsed
}

// Stats returns current pool statistics
func (p *Pool) Stats() PoolStats {
	// Your implementation
	return PoolStats{}
}

// HealthCheck removes unhealthy connections, returns count removed
func (p *Pool) HealthCheck() int {
	// Your implementation
	// - Find unhealthy connections
	// - Remove them from slice
	// - Be careful with slice modification during iteration!
	return 0
}

// Close closes all connections
func (p *Pool) Close() {
	// Your implementation
	// Clear the connections slice
}

// List returns all connections (for display)
func (p *Pool) List() []*Connection {
	return p.connections
}

func main() {
	pool := NewPool("localhost", 5432, 5)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Connection Pool Demo")
	fmt.Println("Type 'help' for commands")
	fmt.Println()

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		cmd := parts[0]

		switch cmd {
		case "help":
			fmt.Println("Commands:")
			fmt.Println("  acquire              - Get a connection")
			fmt.Println("  release <id>         - Return connection")
			fmt.Println("  execute <id> <query> - Run query")
			fmt.Println("  list                 - Show connections")
			fmt.Println("  stats                - Show statistics")
			fmt.Println("  health               - Run health check")
			fmt.Println("  close                - Close all")
			fmt.Println("  quit                 - Exit")

		case "acquire":
			// Your implementation

		case "release":
			// Parse ID and release
			// Your implementation

		case "execute":
			// Parse ID and query, execute
			// Your implementation

		case "list":
			// Your implementation

		case "stats":
			// Your implementation

		case "health":
			// Your implementation

		case "close":
			// Your implementation

		case "quit":
			pool.Close()
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Printf("Unknown command: %s\n", cmd)
		}
	}
}
```

## Test Scenarios

**Scenario 1: Basic lifecycle**
```
acquire        → connection #1
acquire        → connection #2
list           → both IN USE
release 1      → #1 available
acquire        → reuses #1, not #3
```

**Scenario 2: Pool exhaustion**
```
# With maxSize=2:
acquire        → #1
acquire        → #2
acquire        → error: pool exhausted
release 1
acquire        → reuses #1
```

**Scenario 3: Query tracking**
```
acquire
execute 1 SELECT 1
execute 1 SELECT 2
execute 1 SELECT 3
list           → queries: 3
```

**Scenario 4: Health check**
```
# Modify LastUsed to 10 minutes ago (in code)
# Then:
health         → 1 connection removed
```

## Pointer Rules of Thumb

| Situation | Use Pointer? | Why |
|-----------|-------------|-----|
| Method modifies receiver | Yes (*T) | Mutation requires pointer |
| Struct is large | Yes (*T) | Avoid copy overhead |
| Need nil to mean "no value" | Yes (*T) | Values can't be nil |
| Method is read-only, struct is small | No (T) | Copy is safe and clear |
| Slice/map/channel parameter | No | Already reference types |
| Returning new struct | Usually *T | Clearer lifecycle |

## Pitfalls to Avoid

1. **Dereferencing nil**: Always check `if ptr != nil` before `*ptr`
2. **Loop variable address**: `&v` in `for _, v := range` is always same address
3. **Unintended aliasing**: Multiple pointers to same data = shared mutations
4. **Copying pointer receiver structs**: `copy := *ptr` creates independent copy
5. **Mixing receiver types**: Pick one style per type, document why

## Stretch Goals

1. **Connection timeout**: Connections waiting too long in pool get closed
2. **Retry logic**: Acquire retries with backoff when pool is exhausted
3. **Metrics**: Track acquire latency, wait times
4. **Thread safety preview**: Think about what breaks with concurrent access
5. **Connection factory**: Generic pool that works with any resource type

## Time Target

45-60 minutes

## References
- [Go Tour: Pointers](https://go.dev/tour/moretypes/1)
- [Go by Example: Pointers](https://gobyexample.com/pointers)
- [Effective Go: Pointers vs Values](https://go.dev/doc/effective_go#pointers_vs_values)
- [Go FAQ: Methods on Values or Pointers](https://go.dev/doc/faq#methods_on_values_or_pointers)
