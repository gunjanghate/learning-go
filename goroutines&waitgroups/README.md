# Goroutines & WaitGroups

This folder demonstrates how to run concurrent tasks in Go using goroutines and how to wait for them to finish using `sync.WaitGroup`.

## What's Inside

- `main.go`: Shows how to start multiple goroutines, how to track them with a `WaitGroup`, and how to ensure all goroutines complete before `main` exits.
- Focus: Understand basic concurrency with goroutines and synchronization with WaitGroups.

## Key Concepts

### Goroutines

- A **goroutine** is a lightweight thread managed by the Go runtime.
- Use the `go` keyword before a function call to run it concurrently:

  ```go
  go tasks(i, &wg)
  ```

- Goroutines are cheaper than OS threads and are multiplexed onto a smaller number of OS threads by the scheduler.

### WaitGroups

- `sync.WaitGroup` is used to wait for a collection of goroutines to finish.
- Typical pattern:

  ```go
  var wg sync.WaitGroup

  for i := 0; i < 4; i++ {
      wg.Add(1)              // increment counter
      go tasks(i, &wg)       // start goroutine
  }

  wg.Wait()                  // block until counter goes back to 0
  ```

- Inside the goroutine, call `wg.Done()` (usually with `defer`) to signal completion:

  ```go
  func tasks(id int, w *sync.WaitGroup) {
      defer w.Done()
      fmt.Println("Task", id, "is starting")
  }
  ```

## Run The Example

Use PowerShell and run:

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\goroutines&waitgroups"
go run main.go
```

You will see output similar to (order may vary):

```text
Task 3 is starting
Task 2 is starting
Task 0 is starting
Task 1 is starting
```

The order is not guaranteed because the tasks run concurrently.

## Notes and Learning Tips

- Without a `WaitGroup` (or another sync mechanism), `main` might exit before goroutines finish.
- Avoid using `time.Sleep` just to wait for goroutines; `WaitGroup` is the correct tool for this pattern.
- Each goroutine gets its own stack that can grow and shrink, which is why they are lightweight.

## Common `defer` Use Cases

In this example, `defer w.Done()` ensures the WaitGroup is always signaled, even if the function exits early. The same pattern is critical in many real-world scenarios:

| Use Case                                     | Why it's Critical              |
| -------------------------------------------- | ------------------------------ |
| 🔐 `Unlock()` mutex                          | Prevent deadlocks              |
| 🔌 `Close()` DB rows / files / network conns | Prevent resource leaks         |
| 📊 Tracing / metrics                         | Guaranteed logging             |
| 🔄 Cleanup operations                        | Makes code safe + maintainable |
| 🧯 panic recovery                            | Stabilizes API services        |

## Small Example (simplified)

```go
func main() {
    var wg sync.WaitGroup

    for i := 0; i < 4; i++ {
        wg.Add(1)
        go tasks(i, &wg)
    }

    wg.Wait()
}
```

This starts four concurrent tasks and waits for all of them to complete before exiting.
