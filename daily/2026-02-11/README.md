# Day 5: Arrays, Slices, and Slice Operations

**Theme:** Understanding Go's most used data structure

## Learning Objectives
- Understand the difference between arrays (fixed) and slices (dynamic)
- Master slice internals: length, capacity, and the backing array
- Use slice operations: append, copy, slicing expressions
- Avoid common slice pitfalls (aliasing, capacity surprises)
- Know when to pre-allocate with `make()`

## Challenge: Playlist Manager

Build an in-memory playlist manager that demonstrates slice operations through real-world scenarios.

### Core Concepts

```go
// Arrays are fixed-size, value types
var arr [5]int              // [0 0 0 0 0]
arr2 := [...]int{1, 2, 3}   // compiler counts: [3]int

// Slices are dynamic, reference types  
var s []int                  // nil slice
s2 := []int{1, 2, 3}        // literal
s3 := make([]int, 5)        // len=5, cap=5
s4 := make([]int, 0, 10)    // len=0, cap=10 (pre-allocated)
```

### Requirements

1. **Playlist Structure**
   ```go
   type Song struct {
       Title    string
       Artist   string
       Duration int // seconds
   }
   
   type Playlist struct {
       Name  string
       Songs []Song
   }
   ```

2. **Basic Operations**
   - `Add(song Song)` — append to end
   - `Remove(index int)` — remove song at index (preserve order)
   - `RemoveFast(index int)` — remove without preserving order (swap trick)
   - `Insert(index int, song Song)` — insert at position
   - `Get(index int) Song` — retrieve song (bounds checked)

3. **Slice Manipulations**
   - `Shuffle()` — randomize order (Fisher-Yates)
   - `Reverse()` — reverse in place
   - `First(n int) []Song` — return first n songs
   - `Last(n int) []Song` — return last n songs
   - `Slice(start, end int) []Song` — return songs[start:end]

4. **Analysis Functions**
   - `TotalDuration() int` — sum of all durations
   - `Filter(fn func(Song) bool) []Song` — filter by predicate
   - `Map(fn func(Song) string) []string` — transform to strings

5. **Batch Operations**
   - `Merge(other Playlist)` — append another playlist
   - `Dedupe()` — remove duplicate songs (by title+artist)
   - `Split(n int) []Playlist` — split into chunks of n songs

### CLI Interface

```
$ go run main.go
Playlist Manager

Commands:
  add <title> <artist> <duration>  - Add a song
  remove <index>                    - Remove song at index
  insert <index> <title> <artist> <duration>
  list                             - Show all songs
  shuffle                          - Randomize order
  reverse                          - Reverse order  
  first <n>                        - Show first n songs
  last <n>                         - Show last n songs
  duration                         - Total duration
  filter <artist>                  - Show songs by artist
  merge                            - Merge from file
  dedupe                           - Remove duplicates
  debug                            - Show len/cap info
  quit                             - Exit

> add "Bohemian Rhapsody" "Queen" 354
Added: Bohemian Rhapsody by Queen (5:54)

> add "Stairway to Heaven" "Led Zeppelin" 482
Added: Stairway to Heaven by Led Zeppelin (8:02)

> add "Hotel California" "Eagles" 391
Added: Hotel California by Eagles (6:31)

> list
Playlist: My Playlist (3 songs, 20:27 total)
1. Bohemian Rhapsody — Queen (5:54)
2. Stairway to Heaven — Led Zeppelin (8:02)
3. Hotel California — Eagles (6:31)

> debug
Length: 3, Capacity: 4
Backing array can hold 1 more before reallocation

> shuffle
Shuffled!

> list  
1. Hotel California — Eagles (6:31)
2. Bohemian Rhapsody — Queen (5:54)
3. Stairway to Heaven — Led Zeppelin (8:02)

> first 2
1. Hotel California — Eagles (6:31)
2. Bohemian Rhapsody — Queen (5:54)

> insert 1 "Kashmir" "Led Zeppelin" 517
Inserted at position 1

> list
1. Hotel California — Eagles (6:31)
2. Kashmir — Led Zeppelin (8:37)
3. Bohemian Rhapsody — Queen (5:54)
4. Stairway to Heaven — Led Zeppelin (8:02)

> filter "Led Zeppelin"
1. Kashmir — Led Zeppelin (8:37)
2. Stairway to Heaven — Led Zeppelin (8:02)
```

### Key Patterns to Master

**Remove while preserving order:**
```go
func (p *Playlist) Remove(i int) {
    // Shift everything left
    copy(p.Songs[i:], p.Songs[i+1:])
    p.Songs = p.Songs[:len(p.Songs)-1]
}
```

**Remove without preserving order (O(1)):**
```go
func (p *Playlist) RemoveFast(i int) {
    // Swap with last, then shrink
    last := len(p.Songs) - 1
    p.Songs[i] = p.Songs[last]
    p.Songs = p.Songs[:last]
}
```

**Insert at position:**
```go
func (p *Playlist) Insert(i int, song Song) {
    p.Songs = append(p.Songs, Song{}) // grow by one
    copy(p.Songs[i+1:], p.Songs[i:])  // shift right
    p.Songs[i] = song                  // insert
}
```

**Fisher-Yates shuffle:**
```go
func (p *Playlist) Shuffle() {
    for i := len(p.Songs) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        p.Songs[i], p.Songs[j] = p.Songs[j], p.Songs[i]
    }
}
```

**Filter pattern:**
```go
func (p *Playlist) Filter(fn func(Song) bool) []Song {
    // Pre-allocate with zero length but some capacity
    result := make([]Song, 0, len(p.Songs)/2)
    for _, s := range p.Songs {
        if fn(s) {
            result = append(result, s)
        }
    }
    return result
}
```

### Rules
- Use pointer receiver for mutations (`*Playlist`)
- Check bounds before accessing by index
- Use `copy()` appropriately
- Pre-allocate with `make([]T, 0, capacity)` where it makes sense
- stdlib only (`fmt`, `bufio`, `os`, `strings`, `math/rand`)

## Skeleton Code

```go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Song struct {
	Title    string
	Artist   string
	Duration int // seconds
}

func (s Song) String() string {
	min := s.Duration / 60
	sec := s.Duration % 60
	return fmt.Sprintf("%s — %s (%d:%02d)", s.Title, s.Artist, min, sec)
}

type Playlist struct {
	Name  string
	Songs []Song
}

// NewPlaylist creates an empty playlist with pre-allocated capacity
func NewPlaylist(name string) *Playlist {
	return &Playlist{
		Name:  name,
		Songs: make([]Song, 0, 10), // start with capacity 10
	}
}

// Add appends a song to the end
func (p *Playlist) Add(song Song) {
	// Your implementation
}

// Remove removes song at index (preserves order)
func (p *Playlist) Remove(i int) error {
	// Your implementation - return error if out of bounds
	return nil
}

// RemoveFast removes song at index (does NOT preserve order)
func (p *Playlist) RemoveFast(i int) error {
	// Your implementation - swap with last element
	return nil
}

// Insert adds song at specific position
func (p *Playlist) Insert(i int, song Song) error {
	// Your implementation
	return nil
}

// Get retrieves song at index
func (p *Playlist) Get(i int) (Song, error) {
	// Your implementation
	return Song{}, nil
}

// Shuffle randomizes song order using Fisher-Yates
func (p *Playlist) Shuffle() {
	// Your implementation
}

// Reverse reverses the playlist in place
func (p *Playlist) Reverse() {
	// Your implementation
}

// First returns the first n songs (or all if n > len)
func (p *Playlist) First(n int) []Song {
	// Your implementation - use slice expression
	return nil
}

// Last returns the last n songs
func (p *Playlist) Last(n int) []Song {
	// Your implementation
	return nil
}

// TotalDuration returns sum of all song durations
func (p *Playlist) TotalDuration() int {
	// Your implementation
	return 0
}

// Filter returns songs matching the predicate
func (p *Playlist) Filter(fn func(Song) bool) []Song {
	// Your implementation
	return nil
}

// Dedupe removes duplicate songs (same title and artist)
func (p *Playlist) Dedupe() {
	// Your implementation - hint: use a map to track seen
}

// Debug prints length and capacity info
func (p *Playlist) Debug() {
	fmt.Printf("Length: %d, Capacity: %d\n", len(p.Songs), cap(p.Songs))
	remaining := cap(p.Songs) - len(p.Songs)
	if remaining > 0 {
		fmt.Printf("Backing array can hold %d more before reallocation\n", remaining)
	} else {
		fmt.Println("Next append will reallocate!")
	}
}

func (p *Playlist) List() {
	total := p.TotalDuration()
	fmt.Printf("Playlist: %s (%d songs, %d:%02d total)\n", 
		p.Name, len(p.Songs), total/60, total%60)
	for i, s := range p.Songs {
		fmt.Printf("%d. %s\n", i+1, s)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	playlist := NewPlaylist("My Playlist")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Playlist Manager")
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

		// Parse command and arguments
		// Your implementation

		// Handle commands: add, remove, insert, list, shuffle, reverse,
		// first, last, duration, filter, dedupe, debug, help, quit
	}
}
```

## Test Scenarios

**Scenario 1: Basic operations**
```
add "Song A" "Artist 1" 180
add "Song B" "Artist 2" 200
add "Song C" "Artist 1" 220
list
remove 1
list
```

**Scenario 2: Slice aliasing trap**
```
debug            # Check initial capacity
first 2          # This creates a slice sharing backing array!
# Think: what happens if you modify the returned slice?
```

**Scenario 3: Dedupe**
```
add "Same Song" "Same Artist" 100
add "Different" "Artist" 150  
add "Same Song" "Same Artist" 100
list
dedupe
list
```

## Pitfalls to Avoid

1. **Slice aliasing**: `songs[0:3]` shares memory with `songs`
2. **Nil vs empty slice**: `var s []int` vs `s := []int{}`
3. **Capacity surprise**: appending to a slice from `[0:n]` may overwrite
4. **Range copies**: `for _, song := range songs` gives you a copy, not a reference

## Stretch Goals
- Implement `Clone()` that creates a truly independent copy
- Add `Sort(by string)` to sort by title, artist, or duration
- Implement playlist file save/load
- Add undo functionality (track previous states)
- Implement circular iteration (after last song, wrap to first)

## Time Target
45-60 minutes

## References
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Go by Example: Slices](https://gobyexample.com/slices)
- [Slice Tricks Wiki](https://go.dev/wiki/SliceTricks)
- [Effective Go: Slices](https://go.dev/doc/effective_go#slices)
