# Interfaces

This folder demonstrates interfaces in Go using a simple payment gateway example. It shows how to define an interface, implement it with concrete types, and use composition to plug different implementations into a higher-level type.

## What's Inside

- `main.go`: Defines a `paymentr` interface with `pay` and `refund` methods, a `payment` struct that depends on any type implementing `paymentr`, and a concrete implementation `razorpayx`.
- Focus: Understand how interfaces enable polymorphism and decouple code from specific implementations.

## Key Concepts

- **Interface definition**:

  ```go
  type paymentr interface {
      pay(amt float32)
      refund(amt float32, acc string)
  }
  ```

  Any type that provides these two methods with the same signatures implicitly implements `paymentr`.

- **Concrete implementation**:

  ```go
  type razorpayx struct{}

  func (r razorpayx) pay(amt float32) {
      fmt.Println("Payment of", amt, "made using Razorpay")
  }

  func (r razorpayx) refund(amt float32, acc string) {
      fmt.Println("Refund of", amt, "made to account", acc, "using Razorpay")
  }
  ```

- **Struct depending on interface, not concrete type**:

  ```go
  type payment struct {
      gateway paymentr
  }

  func (p payment) pay(amt float32) {
      p.gateway.pay(amt)
  }
  ```

  Here `payment` does not care which payment gateway is used. It only requires something that implements `paymentr`.

- **Polymorphism via interfaces**:
  - You can swap `razorpayx` with another type (e.g., `stripe`, `fakepay`) as long as it implements `paymentr`.
  - This makes your code easier to test and extend.

## Run The Example

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\interfaces"
go run main.go
```

You should see output similar to:

```text
Payment of 500 made using Razorpay
```

## Notes and Learning Tips

- Interfaces in Go are **implicit**: you do not need to "declare" that a type implements an interface; it happens automatically when method sets match.
- Code to interfaces, not implementations: depend on behaviors (methods) instead of concrete types.
- This pattern is very common in real-world Go code for things like logging, storage, HTTP clients, and payment gateways.

## Small Example (simplified)

```go
func main() {
    razor := razorpayx{}

    pay := payment{
        gateway: razor,
    }

    pay.pay(500)
}
```

This creates a `payment` that uses `razorpayx` as its gateway and calls `pay` through the interface.
