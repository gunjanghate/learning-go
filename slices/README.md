# Slices

This folder explores slices in Go: dynamic-size collections built on arrays with convenient methods and flexibility.

## What's Inside

- `main.go`: Examples of slice creation with `make()`, `append()`, slice operators, `copy()`, and built-in slice functions like `slices.Equal()` and `slices.BinarySearch()`.
- Focus: Understanding dynamic sizing, capacity growth, and practical slice operations.

## Key Concepts

- **Dynamic size**: Slices can grow and shrink; size is not fixed at declaration.
- **Backed by arrays**: Slices are views into underlying arrays; they reference, not copy.
- **Capacity vs Length**: `len()` = current elements, `cap()` = underlying array size.
- **Append growth**: When appending beyond capacity, Go allocates a new array (usually double capacity) and copies elements.
- **Slice operator**: `s[start:end]` extracts a subslice from index `start` to `end-1`.
- **Convenient methods**: `append()`, `copy()`, and package `slices` provide useful utilities.

## Run The Example

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\slices"
go run main.go
```

## Notes

- Prefer slices over arrays for most use cases; they're more flexible.
- Understanding capacity helps optimize memory and avoid frequent reallocations.

## Examples

Below are very basic examples to help understand the syntax. The full examples are available in `main.go`.

### Create a Slice

```go
package main

import "fmt"

func main() {
    nums := make([]int, 0, 5)  // slice with length 0 and capacity 5
    fmt.Println(len(nums))     // 0 (length)
    fmt.Println(cap(nums))     // 5 (capacity)
}
```

### Append Elements

```go
package main

import "fmt"

func main() {
    nums := []int{1, 2}
    nums = append(nums, 3)     // append a single value
    nums = append(nums, 4, 5)  // append multiple values
    fmt.Println(nums)          // [1 2 3 4 5]
}
```

### Slice Operator

```go
package main

import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5}
    b := a[1:3]                // from index 1 up to 2
    fmt.Println(b)             // [2 3]
}
```

### Copy Slices

```go
package main

import "fmt"

func main() {
    original := []int{1, 2, 3}
    copied := make([]int, len(original))
    copy(copied, original)
    fmt.Println(copied)        // [1 2 3]
}
```
