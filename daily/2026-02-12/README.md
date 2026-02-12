# Day 6: Maps and Structs

**Theme:** Go's core data modeling tools — composing types for real systems

## Learning Objectives
- Understand map internals: hashing, zero values, nil vs empty maps
- Master map operations: creation, access, deletion, iteration
- Design expressive structs with embedded types
- Use struct tags for metadata
- Combine maps and structs for practical data models

## Challenge: In-Memory Key-Value Store with TTL

Build a Redis-like in-memory cache that uses maps and structs to store data with automatic expiration. This is a pattern you'll use constantly in backend development.

### Core Concepts

```go
// Maps are reference types (nil by default)
var m map[string]int          // nil map — will panic on write!
m2 := make(map[string]int)    // initialized, ready to use
m3 := map[string]int{"a": 1}  // literal initialization

// Comma-ok idiom (crucial!)
val, ok := m["key"]           // ok is false if key doesn't exist
if !ok {
    // handle missing key
}

// Delete is safe on nil maps and missing keys
delete(m, "key")

// Structs group related data
type Entry struct {
    Value     any       // any = interface{}
    CreatedAt time.Time
    TTL       time.Duration
}

// Embedded structs (composition over inheritance)
type TimestampedEntry struct {
    Entry              // embedded — fields promoted
    UpdatedAt time.Time
}
```

### Requirements

1. **Entry Structure**
   ```go
   type Entry struct {
       Value     any
       CreatedAt time.Time
       ExpiresAt time.Time  // zero value = never expires
   }
   
   func (e Entry) IsExpired() bool
   func (e Entry) TTL() time.Duration  // remaining time, or 0 if expired
   ```

2. **Cache Structure**
   ```go
   type Cache struct {
       data     map[string]Entry
       stats    CacheStats
       // mutex for thread safety (stretch goal)
   }
   
   type CacheStats struct {
       Hits       int
       Misses     int
       Sets       int
       Deletes    int
       Evictions  int  // expired entries removed
   }
   ```

3. **Basic Operations**
   - `Set(key string, value any)` — store with no expiration
   - `SetWithTTL(key string, value any, ttl time.Duration)` — store with TTL
   - `Get(key string) (any, bool)` — retrieve value (false if missing or expired)
   - `Delete(key string) bool` — remove entry (returns true if existed)
   - `Exists(key string) bool` — check if key exists and not expired
   - `Keys() []string` — all non-expired keys
   - `Len() int` — count of non-expired entries

4. **TTL Operations**
   - `TTL(key string) (time.Duration, bool)` — time remaining (0 if no TTL)
   - `Expire(key string, ttl time.Duration) bool` — set TTL on existing key
   - `Persist(key string) bool` — remove TTL from key
   - `ExpireAt(key string, t time.Time) bool` — set absolute expiration

5. **Batch Operations**
   - `MSet(pairs map[string]any)` — set multiple keys
   - `MGet(keys []string) map[string]any` — get multiple (only found keys)
   - `Clear()` — delete all entries
   - `Cleanup() int` — remove all expired entries, return count

6. **Stats & Debugging**
   - `Stats() CacheStats` — return cache statistics
   - `Dump() map[string]Entry` — return copy of all data (for debugging)
   - `Info() string` — formatted stats summary

### CLI Interface

```
$ go run main.go
Mini Redis — In-Memory Key-Value Store

Commands:
  set <key> <value> [ttl]    - Store value (ttl in seconds, optional)
  get <key>                  - Retrieve value
  del <key>                  - Delete key
  exists <key>               - Check if key exists
  ttl <key>                  - Get remaining TTL
  expire <key> <seconds>     - Set TTL on existing key
  persist <key>              - Remove TTL
  keys                       - List all keys
  keys <pattern>             - List keys matching pattern (stretch)
  cleanup                    - Remove expired entries
  stats                      - Show cache statistics
  info                       - Show detailed info
  clear                      - Delete everything
  quit                       - Exit

> set name "Stephen"
OK

> set session abc123 300
OK (expires in 5m0s)

> get name
"Stephen"

> get session
"abc123" (TTL: 4m58s)

> ttl name
-1 (no expiration)

> ttl session
298

> expire name 60
OK

> ttl name
59

> persist name
OK

> keys
1) name
2) session

> set temp "will expire" 5
OK (expires in 5s)

> get temp
"will expire" (TTL: 3s)

> # wait 5 seconds...

> get temp
(nil)

> stats
Hits: 4
Misses: 1
Sets: 4
Deletes: 0
Evictions: 1

> info
Mini Redis Stats
================
Keys:        2
Memory:      ~256 bytes (estimated)
Uptime:      2m15s
Hit Rate:    80.0%
```

### Key Patterns to Master

**Safe map access:**
```go
func (c *Cache) Get(key string) (any, bool) {
    entry, exists := c.data[key]
    if !exists {
        c.stats.Misses++
        return nil, false
    }
    if entry.IsExpired() {
        delete(c.data, key)  // lazy eviction
        c.stats.Evictions++
        c.stats.Misses++
        return nil, false
    }
    c.stats.Hits++
    return entry.Value, true
}
```

**Map iteration (order is random!):**
```go
func (c *Cache) Keys() []string {
    keys := make([]string, 0, len(c.data))
    for k, entry := range c.data {
        if !entry.IsExpired() {
            keys = append(keys, k)
        }
    }
    return keys
}
```

**Struct method with pointer receiver:**
```go
func (c *Cache) Set(key string, value any) {
    c.data[key] = Entry{
        Value:     value,
        CreatedAt: time.Now(),
        // ExpiresAt zero value = never
    }
    c.stats.Sets++
}
```

**Time comparison:**
```go
func (e Entry) IsExpired() bool {
    if e.ExpiresAt.IsZero() {
        return false  // no expiration set
    }
    return time.Now().After(e.ExpiresAt)
}
```

### Rules
- Always use comma-ok idiom for map reads
- Never write to a nil map
- Use pointer receivers for methods that modify state
- Handle expiration lazily (on access) or via periodic cleanup
- stdlib only (`fmt`, `bufio`, `os`, `strings`, `time`, `strconv`)

## Skeleton Code

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Value     any
	CreatedAt time.Time
	ExpiresAt time.Time // zero = never expires
}

// IsExpired returns true if the entry has passed its expiration time
func (e Entry) IsExpired() bool {
	// Your implementation
	return false
}

// TTL returns remaining time to live, or 0 if expired/no TTL
func (e Entry) TTL() time.Duration {
	// Your implementation
	return 0
}

type CacheStats struct {
	Hits      int
	Misses    int
	Sets      int
	Deletes   int
	Evictions int
}

type Cache struct {
	data      map[string]Entry
	stats     CacheStats
	createdAt time.Time
}

// NewCache creates an initialized cache
func NewCache() *Cache {
	return &Cache{
		data:      make(map[string]Entry),
		createdAt: time.Now(),
	}
}

// Set stores a value with no expiration
func (c *Cache) Set(key string, value any) {
	// Your implementation
}

// SetWithTTL stores a value that expires after ttl duration
func (c *Cache) SetWithTTL(key string, value any, ttl time.Duration) {
	// Your implementation
}

// Get retrieves a value, returning false if missing or expired
func (c *Cache) Get(key string) (any, bool) {
	// Your implementation — remember lazy eviction!
	return nil, false
}

// Delete removes a key, returning true if it existed
func (c *Cache) Delete(key string) bool {
	// Your implementation
	return false
}

// Exists checks if key exists and is not expired
func (c *Cache) Exists(key string) bool {
	// Your implementation
	return false
}

// Keys returns all non-expired keys
func (c *Cache) Keys() []string {
	// Your implementation
	return nil
}

// Len returns count of non-expired entries
func (c *Cache) Len() int {
	// Your implementation
	return 0
}

// TTL returns remaining time for a key (-1 = no expiration, 0 = expired/missing)
func (c *Cache) TTL(key string) (time.Duration, bool) {
	// Your implementation
	return 0, false
}

// Expire sets a TTL on an existing key
func (c *Cache) Expire(key string, ttl time.Duration) bool {
	// Your implementation
	return false
}

// Persist removes expiration from a key
func (c *Cache) Persist(key string) bool {
	// Your implementation
	return false
}

// MSet sets multiple key-value pairs
func (c *Cache) MSet(pairs map[string]any) {
	// Your implementation
}

// MGet retrieves multiple keys (only returns found, non-expired)
func (c *Cache) MGet(keys []string) map[string]any {
	// Your implementation
	return nil
}

// Cleanup removes all expired entries and returns count removed
func (c *Cache) Cleanup() int {
	// Your implementation
	return 0
}

// Clear removes all entries
func (c *Cache) Clear() {
	// Your implementation
}

// Stats returns cache statistics
func (c *Cache) Stats() CacheStats {
	return c.stats
}

// Info returns a formatted stats summary
func (c *Cache) Info() string {
	// Your implementation — include uptime, key count, hit rate
	return ""
}

func main() {
	cache := NewCache()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Mini Redis — In-Memory Key-Value Store")
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

		// Parse and handle commands
		// Your implementation
	}
}
```

## Test Scenarios

**Scenario 1: Basic CRUD**
```
set user "Stephen"
get user
exists user
del user
get user
exists user
```

**Scenario 2: TTL behavior**
```
set token "abc123" 3
get token
ttl token
# wait 4 seconds
get token
stats
```

**Scenario 3: Modify existing TTL**
```
set data "important"
ttl data
expire data 60
ttl data
persist data
ttl data
```

**Scenario 4: Bulk operations**
```
# In code: cache.MSet(map[string]any{"a": 1, "b": 2, "c": 3})
keys
# In code: cache.MGet([]string{"a", "b", "x"})  // x doesn't exist
```

## Pitfalls to Avoid

1. **Writing to nil map**: Always use `make()` or literal initialization
2. **Forgetting comma-ok**: `val := m["key"]` returns zero value if missing
3. **Map iteration order**: Never rely on map order — it's randomized
4. **Concurrent map access**: Maps are NOT thread-safe (stretch goal adds mutex)
5. **Struct field visibility**: lowercase = unexported (private to package)

## Stretch Goals

1. **Pattern matching**: Implement `Keys(pattern string)` with wildcard support
2. **Thread safety**: Add `sync.RWMutex` for concurrent access
3. **Background cleanup**: Goroutine that periodically evicts expired entries
4. **Memory estimation**: Calculate approximate memory usage
5. **Type-specific commands**: `Incr(key)`, `Append(key, str)`, `LPush(key, val)`
6. **Persistence**: Save/load cache to JSON file

## Time Target

45-60 minutes

## References
- [Go Maps in Action](https://go.dev/blog/maps)
- [Go by Example: Maps](https://gobyexample.com/maps)
- [Go by Example: Structs](https://gobyexample.com/structs)
- [Effective Go: Maps](https://go.dev/doc/effective_go#maps)
- [Redis Commands Reference](https://redis.io/commands/) (for inspiration)
