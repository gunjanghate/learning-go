# Functions & Variadic Functions

This folder contains examples showing how to declare and use functions in Go, including returning multiple values, passing and returning functions, and variadic functions.

## What's Inside

- `main.go`: Demonstrates function declaration, multiple return values, higher-order functions (functions that return other functions), and variadic functions using `...`.
- Focus: Learn how to structure functions, use return values, and work with flexible argument lists.

## Key Concepts

- **Function declaration**: `func name(params) returnType {}`
- **Multiple return values**: Functions can return multiple values, e.g., `(string, string, string)`.
- **Higher-order functions**: Functions can accept other functions as parameters or return functions.
- **Variadic functions**: Use `...T` to accept any number of arguments of type `T`; inside the function you receive a slice of `T`.

## Run The Example

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\functions&variadicfunctions"
go run main.go
```

## Notes

- Use clear parameter names and return variable names when helpful for readability.
- Variadic functions are useful for operations like `sum`, `join`, or formatting where the number of arguments varies.

## Examples

### Simple Function

```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    fmt.Println(add(2,3)) // 5
}
```

### Multiple Return Values

```go
package main

import "fmt"

func getLangs() (string, string, string) {
    return "golang", "python", "java"
}

func main() {
    l1, l2, l3 := getLangs()
    fmt.Println(l1, l2, l3)
}
```

### Higher-order Function (returning a function)

```go
package main

import "fmt"

func processIt() func(a int) int {
    return func(a int) int {
        return a
    }
}

func main() {
    fn := processIt()
    fmt.Println(fn(10)) // 10
}
```

### Variadic Function

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    nums := []int{1,2,3,4,5}
    fmt.Println(sum(nums...)) // 15
}
```
