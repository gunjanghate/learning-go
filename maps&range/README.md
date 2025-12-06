# Maps & Range

This folder demonstrates maps (key-value collections) and the `range` keyword for iteration in Go.

## What's Inside

- `main.go`: Examples showing map creation, insertion, deletion, lookup with presence test, map equality (via helper), and `range` iteration over slices, maps, and strings.
- Focus: How to use maps, handle missing keys, delete entries, and iterate with `range`.

## Key Concepts

- **Maps**: Unordered key-value collections with dynamic size. Keys must be unique and can be of many comparable types; values can be any type.
- **Zero value on missing key**: Accessing a non-existent key returns the zero value of the value type (e.g., `0` for `int`, `""` for `string`). Use the `value, ok := m[key]` pattern to check presence.
- **Delete**: `delete(m, key)` removes an entry.
- **Range**: `for i, v := range collection {}` iterates over slices (index,value), maps (key,value), and strings (index,rune).
- **Strings & runes**: `range` on strings returns the starting byte index and the Unicode code point (rune).

## Run The Example

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\maps&range"
go run main.go
```

## Notes

- Maps are best for lookups by key and when you need dynamic, associative storage.
- Iteration order over maps is not guaranteed; do not rely on ordering.
- Use `range` for concise iteration over common collections.

## Examples

Below are very basic examples to help understand the syntax. The full examples are available in `main.go`.

### Basic Map

```go
package main

import "fmt"

func main() {
    m := make(map[string]string)
    m["name"] = "GG"
    m["course"] = "GoLang"

    fmt.Println("Name:", m["name"])   // GG
    fmt.Println("Age:", m["age"])     // "" (zero value, key missing)

    // presence check
    v, ok := m["age"]
    fmt.Println(v, ok) // "" false
}
```

Expected output:

```
Name: GG
Age:

"" false
```

### Delete & Length

```go
package main

import "fmt"

func main() {
    m := map[string]string{"a":"1", "b":"2"}
    delete(m, "a")
    fmt.Println(len(m)) // 1
}
```

Expected output: `1`

### Range iteration

```go
package main

import "fmt"

func main() {
    nums := []int{6, 7, 8}
    for i, v := range nums {
        fmt.Println(i, v) // index and value
    }

    scores := map[string]int{"maths": 90, "english": 88}
    for k, v := range scores {
        fmt.Println(k, v) // key and value (order not guaranteed)
    }

    s := "hello"
    for i, r := range s {
        fmt.Println(i, r) // index (byte offset) and rune value
    }
}
```

Expected output (order for map may vary):

```
0 6
1 7
2 8
maths 90
english 88
0 104
1 101
2 108
3 108
4 111
```
