# Email Processing Assignment

This folder contains a concurrency assignment that builds an email processing system using goroutines, channels, WaitGroups, and a simple global rate limiter.

The goal is to practice:

- Designing worker-style concurrent systems
- Using channels safely
- Coordinating shutdown with `WaitGroup`
- Applying rate limiting with time-based signals

## Files

- `emailprocessing.go`: Main program implementing the email job producer and worker pool.

## Requirements (Problem Statement)

From the comments at the top of `emailprocessing.go`, the original requirements are:

1. **Dispatcher goroutine**

   - Reads incoming email jobs from a main job channel.
   - Assigns jobs to a pool of worker goroutines.

2. **Worker pool (N workers)**

   - Each worker:
     - Receives jobs via a channel.
     - Simulates sending email (e.g. with `time.Sleep`).
     - Prints: `"Worker X sent email to <address>"`.

3. **Use `WaitGroup`** to wait for **all workers** to finish.

4. **Use buffered channels** for job queue.

5. **Support graceful shutdown**
   - When `main` closes the job channel:
     - Dispatcher must not panic.
     - Workers must stop after finishing their current tasks.
     - `main` waits for all workers to exit.

## Current Implementation Overview

The active code in `emailprocessing.go` focuses on:

- A shared job queue (`jobQueue`) for email jobs.
- A worker pool of goroutines consuming from that queue.
- A **global rate limiter** using `time.Tick`.
- Proper shutdown using `close(jobQueue)` and `WaitGroup`.

### Data Model: `EmailJob`

```go
type EmailJob struct {
    ID    int
    Email string
}
```

Each job has a numeric `ID` and an `Email` address.

### Worker Function

```go
func worker(id int, jobQueue <-chan EmailJob, wg *sync.WaitGroup, limiter <-chan time.Time) {
    defer wg.Done()
    for job := range jobQueue {
        <-limiter // global rate limit before processing

        fmt.Println("Worker", id, "sending email to:", job.Email)
        time.Sleep(time.Second) // simulate slow email API
    }
    fmt.Println("Worker", id, "exiting...")
}
```

Key points:

- **Channel directions**:
  - `jobQueue <-chan EmailJob`: receive-only inside the worker.
  - `limiter <-chan time.Time`: receive-only channel used for rate limiting.
- `defer wg.Done()` ensures that the `WaitGroup` is always signaled when a worker exits.
- The worker loops over `jobQueue` using `range`:
  - When `jobQueue` is closed and drained, the loop ends and the worker prints an exit message.

### Main Function Flow

```go
func main() {
    jobQueue := make(chan EmailJob, 20)
    numWorkers := 3

    // GLOBAL rate limiter: allow 1 email / 300ms => ~3.3 emails per second
    limiter := time.Tick(300 * time.Millisecond)
    var wg sync.WaitGroup

    // Start workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go worker(i, jobQueue, &wg, limiter)
    }

    // Produce jobs
    for i := 0; i < 10; i++ {
        jobQueue <- EmailJob{
            ID:    i,
            Email: fmt.Sprintf("user%d@example.com", i),
        }
    }

    close(jobQueue)
    wg.Wait()
    fmt.Println("All jobs processed!")
}
```

What this does:

- Creates a **buffered** channel `jobQueue` with capacity `20` to hold pending email jobs.
- Configures a global rate limiter using `time.Tick(300 * time.Millisecond)` to allow about 3.3 emails per second, shared by all workers.
- Starts `numWorkers` worker goroutines, all reading from the same `jobQueue`.
- Produces 10 email jobs and sends them into `jobQueue`.
- Calls `close(jobQueue)` to signal that no more jobs will arrive.
- Waits for all workers to finish using `wg.Wait()`.

This design satisfies the key ideas:

- Worker pool over a shared job queue.
- Buffered channel for the job queue.
- Graceful shutdown through channel closing and `WaitGroup` synchronization.
- Simple global rate limiting.

## How to Run

From PowerShell:

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\assignements"
go run .\emailprocessing.go
```

Expected-style output (order and timing will vary):

```text
Worker 0 sending email to: user0@example.com
Worker 1 sending email to: user1@example.com
Worker 2 sending email to: user2@example.com
...
Worker 1 exiting...
Worker 0 exiting...
Worker 2 exiting...
All jobs processed!
```

The exact order depends on the scheduler and timing.

## Concepts Practiced

- **Goroutines**: multiple workers running concurrently.
- **Channels**: communication mechanism between producer (`main`) and consumers (workers).
- **Buffered channels**: smooth out spikes in job production.
- **WaitGroups**: wait for all worker goroutines to complete before exiting.
- **Rate limiting**: use `time.Tick` to control how frequently work is done across all workers.
- **Graceful shutdown**:
  - Closing the job channel when all jobs are enqueued.
  - Ranging over the channel in workers so they exit cleanly when the channel is closed.

## Possible Extensions

If you want to extend this assignment further:

- Reintroduce a **separate dispatcher** goroutine:
  - `main` sends jobs to a dispatcher.
  - Dispatcher fans out jobs to multiple worker channels.
- Add **error handling** and retry logic for failed email sends.
- Add **context cancellation** (using `context.Context`) to support timeouts or manual cancellation.
- Add **metrics/logging** (e.g., count emails sent per worker, measure latency).
- Make the rate limit configurable (e.g., via flags or environment variables).

These changes will bring the assignment closer to real-world email or job processing systems.
