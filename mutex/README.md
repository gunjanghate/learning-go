# Mutex

This folder demonstrates how to use a mutex (`sync.Mutex`) to safely update shared data from multiple goroutines and avoid race conditions.

## What's Inside

- `main.go`: Example where multiple goroutines increment a shared `views` counter on a `post` struct.
- Focus: Understanding race conditions and how mutexes protect shared state.

## Key Concepts

- **Race condition**: Happens when two or more goroutines access the same memory at the same time and at least one of them writes to it. This can lead to incorrect or unpredictable results.
- **Mutex (`sync.Mutex`)**: A mutual exclusion lock. Only one goroutine can hold the lock at a time.
  - `Lock()`: acquire the lock before accessing shared data.
  - `Unlock()`: release the lock after you are done.
- **WaitGroup (`sync.WaitGroup`)**: Used to wait for all goroutines to finish before exiting `main`.

## Code Overview

### `post` struct with embedded mutex

```go
type post struct {
    views int
    mu    sync.Mutex
}
```

- `views`: shared counter we want to protect.
- `mu`: a mutex to guard `views`.

### Increment Method

```go
func (p *post) inc(wg *sync.WaitGroup, i int) {
    p.mu.Lock()
    p.views++
    fmt.Println("Incremented views to:", p.views)
    fmt.Println("Go routine no:", i)
    defer func() {
        wg.Done()
        p.mu.Unlock()
    }()
}
```

- Method has a pointer receiver `*post` so it can modify the original `post`.
- `p.mu.Lock()` ensures that only one goroutine at a time can increment `views`.
- `defer` ensures `wg.Done()` and `p.mu.Unlock()` are always called when the function returns.

### Spawning Goroutines in `main`

```go
func main() {
    var wg sync.WaitGroup
    myPost := post{views: 0}

    for i := range 10 {
        wg.Add(1)
        go myPost.inc(&wg, i)
    }

    wg.Wait()

    fmt.Println("Views:", myPost.views)
}
```

- Creates a `post` with `views` set to `0`.
- Starts 10 goroutines, each calling `myPost.inc(&wg, i)`.
- Uses `wg.Add(1)` before each goroutine and `wg.Wait()` to block until all increments finish.
- Final print shows the total number of views after all increments.

Without the mutex, you would likely see a smaller or inconsistent final `views` value because of concurrent writes.

## Run The Example

From PowerShell:

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\mutex"
go run main.go
```

You will see multiple lines showing the incremented views and goroutine numbers, followed by the final total:

```text
Incremented views to: 1
Go routine no:  0
...
Views: 10
```

The order of goroutine logs may vary, but the final count should be correct thanks to the mutex.
