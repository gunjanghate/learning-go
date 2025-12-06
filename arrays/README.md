# Arrays

This folder explores arrays in Go: fixed-size collections with constant-time access and memory efficiency.

## What's Inside

- `main.go`: Examples of array declaration, initialization, zero values, multi-dimensional arrays, and iteration.
- Focus: Understanding fixed-size collections, memory layout, and when to use arrays vs slices.

## Key Concepts

- **Fixed size**: Arrays have a fixed length defined at declaration; cannot grow or shrink.
- **Same data type**: All elements must be the same type (e.g., `int`, `string`).
- **Constant-time access**: Access any element in O(1) time using an index.
- **Memory-efficient**: Contiguous memory storage for fast access and cache locality.
- **Zero values**: Uninitialized elements get default values (`0` for numbers, `false` for bools, `""` for strings).

## Run The Example

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\arrays"
go run main.go
```

## Notes

- Use arrays when you know the size ahead of time.
- For dynamic collections, prefer slices.

## Examples

Below are very basic examples to help understand the syntax. The full examples are available in `main.go`.

### Basic Array Declaration

```go
package main

import "fmt"

func main() {
    var nums [4]int         // 4-element integer array, initialized to zero values
    fmt.Println(nums)       // [0 0 0 0]

    nums[0] = 10            // assign first element
    nums[1] = 20
    fmt.Println(nums)       // [10 20 0 0]
}
```

### Array with Initialization

```go
package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}  // 3-element array initialized with values
    fmt.Println(arr)        // [1 2 3]
    fmt.Println(len(arr))   // 3 (length)
}
```

### Multi-dimensional Array (2D array)

```go
package main

import "fmt"

func main() {
    matrix := [2][2]int{
        {1, 2},
        {3, 4},
    }
    fmt.Println(matrix)     // [[1 2] [3 4]]
}
```
