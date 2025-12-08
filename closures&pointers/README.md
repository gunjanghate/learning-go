# Closures & Pointers

This folder demonstrates closures and pointers in Go, showing how inner functions capture outer variables (closures) and how pointers allow functions to modify variables by reference.

## What's Inside

- `main.go`: Examples of a closure-producing function (`counter()`), passing by value vs passing by reference (`changeNum` vs `changeNumByRef`), and usage patterns.
- Focus: Understand lexical closures, function lifetime implications, and pointer usage for mutable operations.

## Key Concepts

- **Closures**: A closure is an inner function that references variables from its outer function. Those outer variables remain alive as long as the inner function is reachable.
- **Each closure instance owns its state**: Calling the outer function multiple times produces independent closure instances with their own captured variables.
- **Pointers**: Use `*T` to declare pointer types and `&variable` to get an address. Passing a pointer lets a function modify the original value.
- **Pass by value**: Regular parameters are passed by value (copies are made). To modify the original, pass a pointer.

## Run The Example

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\closures&pointers"
go run main.go
```

## Notes

- Closures are useful for building small stateful functions (e.g., counters, generators).
- Use pointers when you need to update caller state or avoid copying large structures.

## Examples

### Closure (counter)

```go
package main

import "fmt"

func counter() func() int {
    var cnt int = 0
    return func() int {
        cnt++
        return cnt
    }
}

func main() {
    count := counter()
    fmt.Println(count()) // 1
    fmt.Println(count()) // 2
    fmt.Println(count()) // 3
}
```

### Passing by Value vs Reference

```go
package main

import "fmt"

func changeNum(num int) {     // passed by value
    num = 20
    fmt.Println("In changeNum:", num)
}

func changeNumByRef(num *int) { // passed by reference (pointer)
    *num = 30
    fmt.Println("In changeNumByRef:", *num)
}

func main() {
    num := 1
    fmt.Println("Before changeNum:", num) // 1
    changeNum(num)
    fmt.Println("After changeNum:", num)  // still 1

    fmt.Println("Before changeNumByRef:", num) // 1
    changeNumByRef(&num)
    fmt.Println("After changeNumByRef:", num)  // 30
}
```
