# Concurrency Patterns

This folder demonstrates two common concurrency patterns in Go using goroutines, channels, and WaitGroups:

- Fan-out / Fan-in
- Worker Pool

It also includes a diagram to visualize how jobs and workers interact.

![Concurrency Patterns Diagram](diagrams.png)

## What's Inside

- `main.go`:
  - Commented example of **Fan-out / Fan-in** using multiple workers processing image URLs and a single results channel.
  - Active implementation of a **Worker Pool** that processes many image URLs with a fixed number of workers.

## Key Types

```go
type Result struct {
    Value string
    Err   error
}
```

- `Result` wraps the outcome of processing a job (here: processed image URL).

## Pattern 1: Fan-out / Fan-in (in comments)

**Idea:**

- **Fan-out**: Start multiple goroutines, each doing the same kind of work (e.g., processing images).
- **Fan-in**: Collect all results into a single channel and handle them in one place.

In `main.go` (commented section):

- A `WaitGroup` tracks two `worker` goroutines.
- Each goroutine processes one image URL and sends a `Result` into `resChan`.
- After `wg.Wait()` and `close(resChan)`, a loop `range resChan` reads all results (fan-in).

This pattern is useful when:

- You have independent tasks that can run in parallel.
- You want to centralize result handling (logging, aggregation, etc.).

## Pattern 2: Worker Pool (active code)

**Problem:**

- You may have thousands or more jobs (e.g., many images to process).
- Spawning one goroutine per job can exhaust resources.

**Solution:**

- Create a **fixed number of workers** (goroutines).
- Send all jobs into a **jobs channel**.
- Each worker pulls from the channel, processes jobs, and sends results to a **results channel**.

### Job List

```go
jobs := []string{
    "http://image1.com",
    "http://image2.com",
    // ... up to image20
}
```

### Worker Function

```go
func worker(jobsChan chan string, wg *sync.WaitGroup, resChan chan Result) {
    defer wg.Done()
    for job := range jobsChan {
        time.Sleep(50 * time.Millisecond)
        resChan <- Result{
            Value: "Processed " + job,
            Err:   nil,
        }
    }
    fmt.Println("Worker Shutting down")
}
```

- Reads from `jobsChan` in a loop (`range jobsChan`).
- Simulates processing with `time.Sleep`.
- Sends a `Result` into `resChan`.
- Exits cleanly when `jobsChan` is closed.

### Setting Up the Pool

```go
var wg sync.WaitGroup
resChan := make(chan Result, 50)
jobsChan := make(chan string, len(jobs))
const totalWorkers = 5

for i := 1; i <= totalWorkers; i++ {
    wg.Add(1)
    go worker(jobsChan, &wg, resChan)
}
```

- `jobsChan` is buffered to hold all jobs.
- Five worker goroutines share the same `jobsChan`.

### Coordinating Shutdown

```go
stTime := time.Now()

go func() {
    wg.Wait()
    close(resChan)
}()

for _, j := range jobs {
    jobsChan <- j
}
close(jobsChan)

for res := range resChan {
    fmt.Println("Job Completed:", res.Value)
}

fmt.Println("Time taken:", time.Since(stTime))
```

- A goroutine waits for all workers to finish (`wg.Wait()`) and then closes `resChan`.
- Main goroutine:
  - Sends all jobs into `jobsChan`.
  - Closes `jobsChan` to signal no more jobs.
  - Ranges over `resChan` to receive every `Result` until it is closed.
- Finally prints total time taken.

## Run The Example

From PowerShell:

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\concurrencypatterns"
go run main.go
```

You will see lines like:

```text
Job Completed: Processed http://image3.com
...
Worker Shutting down
Time taken: 50.1234ms
```

(Exact order and timing will vary.)

## Notes and Learning Tips

- **Fan-out / Fan-in** is great when each job is simple and independent; you can scale the number of goroutines easily.
- **Worker Pool** is better when you have many jobs and need to **limit concurrency** to avoid overloading CPU, memory, or external services.
- Always:
  - Close job channels from the **sender** when done.
  - Use `WaitGroup` to know when workers are finished.
  - Close result channels from the goroutine that knows when all workers are done (after `wg.Wait()`).

Use the `diagrams.png` image in this folder to visually connect these patterns with the code.
