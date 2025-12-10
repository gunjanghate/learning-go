# Enums & Generics

This folder demonstrates how to model enums using typed constants in Go and how to use generics (type parameters) for reusable functions and data structures.

## What's Inside

- `main.go`: Shows an `OrderStatus` enum-like type, a generic function `printSlice`, and a generic `Stack` type.
- Focus: Learn how to represent enum values, add type safety, and write generic code that works with different types.

## Key Concepts

### Enum-like Types with Constants

Go does not have a built-in `enum` keyword, but you can create enum-like types using custom types and constants.

```go
type OrderStatus string

const (
    Rece  OrderStatus = "Received"
    Conf  OrderStatus = "Confirmed"
    Prep  OrderStatus = "Prepared"
    Deliv OrderStatus = "Delivered"
)

func changeOrderStatus(status OrderStatus) {
    fmt.Println("Changing order status to", status)
}
```

- `OrderStatus` is a custom type based on `string`.
- The `const` block defines the allowed values for this type.
- Using a specific type instead of plain `string` gives better type safety.

> Note: You can also use `iota` for numeric enums; the comments in `main.go` show this idea.

### Generic Functions

Generics allow you to write functions that work with many types while still being type safe.

```go
func printSlice[T comparable, V string](items []T, name V) {
    for _, item := range items {
        fmt.Println(item, name)
    }
}
```

- `T` and `V` are type parameters.
- `T comparable` means `T` must support `==` and `!=` (e.g., `int`, `string`, `bool`).
- `V string` means `V` must be exactly `string`.
- Inside the function, `items` is a slice of `T` and `name` is a `string`.

Example usage from `main`:

```go
ints := []int{1, 2, 3, 4, 5}
printSlice(ints, "GG")
```

### Generic Types (Stack Example)

You can also define generic types using type parameters.

```go
type Stack[T any] struct {
    elements []T
}
```

- `T any` means `T` can be any type.
- `elements` holds a slice of `T`.

Example usage:

```go
st := Stack[int]{
    elements: []int{10, 20, 30, 40, 50},
}

for _, num := range st.elements {
    fmt.Println(num)
}
```

This creates a `Stack[int]` and prints each element.

## Run The Example

Use PowerShell and run:

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\enums&generics"
go run main.go
```

You will see output showing:

- The order status being changed (e.g., `Changing order status to Confirmed`).
- The items printed by `printSlice`.
- The elements from the generic `Stack[int]`.

## Notes and Learning Tips

- Use a custom type + constants pattern to create clear, type-safe enums.
- Start with simple generic functions like `printSlice` or `sum` to build intuition.
- Generic types like `Stack[T]` are useful for reusable data structures (stacks, queues, caches) without losing type safety.
