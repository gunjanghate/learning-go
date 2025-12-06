# Loops & Conditions

This folder explores control flow in Go: loops for iteration and conditionals for decision-making.

## What's Inside

- `main.go`: Examples of `for` loops (while-style, classic, range), `if-else` statements, `switch` cases (simple, multiple conditions, type switch).
- Focus: Understanding Go's unified looping construct and branching patterns.

## Key Concepts

### Loops

- **`for` is the only loop construct**: Go has no `while` or `do-while`; use `for` for all looping.
- **While-style**: `for condition { }` loops while condition is true.
- **Classic for**: `for init; condition; increment { }` works like C-style loops.
- **Range**: `for i := range n { }` iterates from 0 to n-1 (Go 1.22+).
- **Infinite loop**: `for { }` runs forever unless broken.

### Conditionals

- **If-Else**: Standard branching with optional variable declaration in condition.
- **If-Else If**: Chain multiple conditions; only first true block executes.
- **Switch-Case**: Cleaner than many if-else; no break needed (auto-break by default).
- **Multiple conditions**: `case a, b, c:` matches any of these values.
- **Type switch**: `switch v := x.(type) { }` branches on dynamic type.

## Run The Example

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\loops&conditions"
go run main.go
```

## Notes

- Use `for` uniformly; it's simpler than having multiple loop keywords.
- Switch is often clearer than deeply nested if-else.

## Examples

Below are basic examples to help understand the syntax. The full examples are available in `main.go`.

### While-Style Loop

```go
package main

import "fmt"

func main() {
    i := 1
    for i <= 5 {           // loop while condition is true
        fmt.Println(i)     // print i
        i++                // increment i
    }
}
```

Output: `1 2 3 4 5`

### Classic For Loop

```go
package main

import "fmt"

func main() {
    for j := 1; j <= 5; j++ {  // init; condition; post
        fmt.Println(j)
    }
}
```

Output: `1 2 3 4 5`

### If-Else

```go
package main

import "fmt"

func main() {
    num := 10
    if num%2 == 0 {
        fmt.Println("Even")    // even number
    } else {
        fmt.Println("Odd")     // odd number
    }
}
```

Output: `Even`

### Switch-Case

```go
package main

import "fmt"

func main() {
    a := 5
    switch a {
    case 1:
        fmt.Println("One")
    case 5:
        fmt.Println("Five")    // this will print
    default:
        fmt.Println("Other")
    }
}
```

Output: `Five`
