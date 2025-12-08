# Structs & Embedding

This folder demonstrates struct types in Go, methods (including pointer receivers), factory-style constructors, anonymous structs, and embedding one struct inside another.

## What's Inside

- `main.go`: Defines `customer` and `order` structs, shows how to create instances, how to attach methods with pointer receivers to modify struct state, how to use a constructor-like `newOrder` function that returns `*order`, and how to embed a `customer` inside an `order`.
- Focus: Learn how to model data with structs, mutate struct values using methods and pointers, and compose structs by embedding.

## Key Concepts

- **Structs**: Composite types that group fields. Example:

  ```go
  type customer struct {
      name  string
      email string
  }

  type order struct {
      id        string
      amt       float32
      status    string
      createdAt time.Time
      customer  customer // embedding a customer inside an order
  }
  ```

- **Methods and receivers**:

  - A method with a pointer receiver (e.g., `func (o *order) changeStatus(...)`) can modify the original struct instance.
  - Methods with value receivers operate on a copy.

- **Constructor / factory function**:

  - Use a function like `newOrder(...) *order` to initialize and return a pointer to a struct. This keeps initialization logic in one place and returns a modifiable instance.

- **Embedding / composition**:

  - Include one struct as a field of another (e.g., `customer` inside `order`) to compose data models.

- **Anonymous structs**:

  - Useful for small, one-off grouped data without declaring a named type.

- **Struct comparison**:
  - Two structs can be compared with `==` only when all their fields are comparable. The example code has a commented comparison to demonstrate this concept.

## Run The Example

Make sure you are in the folder and run the example using PowerShell:

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\structs&embedding"
go run main.go
```

## Notes and Learning Tips

- Prefer pointer receivers for methods that modify the struct (e.g., update status, amount) or when you want to avoid copying large structs.
- Use factory functions to encapsulate initialization details (timestamps, default values).
- Embedding a struct models a "has-a" relationship and keeps related data together (e.g., an `order` has a `customer`).
- Anonymous structs are convenient for quick examples or local grouping of fields.

## Example Walkthrough (what `main.go` demonstrates)

- Create an `order` with `newOrder(...)`, modify its amount, and call a pointer-receiver method `changeStatus` to update status.
- Create a second `order` with inline struct literal and call `changeAmt` on it.
- Print the full `order` values and specific fields (e.g., `newOrder.customer.email`).
- Demonstrates how struct fields are automatically dereferenced when using pointer receivers in methods.

### Small example (extracted)

```go
o := newOrder("ord001", 250.75, "pending")
o.changeStatus("shipped")

newOrder := order{
    id: "ord003",
    amt: 750.50,
    status: "processing",
    customer: customer{
        name: "John Doe",
        email: "john.doe@example.com",
    },
}
fmt.Println("New Order with Customer:", newOrder)
fmt.Println("Customer details:", newOrder.customer.email)
```

Expected (example) output will show the order structs and the customer's email when you run `main.go`.

---
