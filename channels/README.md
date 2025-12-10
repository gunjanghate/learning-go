# Channels

This folder demonstrates Go channels: how goroutines communicate, the difference between unbuffered and buffered channels, direction-only channels, and using `select` to wait on multiple channels.

## What's Inside

- `main.go`: Examples of
  - Creating channels with `make` (unbuffered and buffered)
  - Using channels to synchronize goroutines (`done` channel)
  - Sending and receiving values on a buffered channel (`emailChan`)
  - Direction-only channels in function signatures
  - Using `select` to read from multiple channels (`chan1`, `chan2`)
- Focus: Understand how channels coordinate work between goroutines.

## Key Concepts

### What is a Channel?

- A channel is a typed conduit through which goroutines can send and receive values.
- Think of it as a pipe between goroutines:

  ```go
  ch := make(chan int)     // unbuffered channel of ints
  chBuf := make(chan int, 5) // buffered channel with capacity 5
  ```

### Unbuffered vs Buffered

- **Unbuffered channel**:

  - Send (`ch <- v`) blocks until another goroutine is ready to receive.
  - Receive (`v := <-ch`) blocks until there is a value to take.
  - Useful for synchronization (ensuring hand-off between sender and receiver).

- **Buffered channel**:

  - Has capacity; send does not block until the buffer is full.
  - Receive blocks only when the buffer is empty.
  - Example from code:

  ```go
  emailChan := make(chan string, 10)
  ```

  This allows you to queue several emails before they are processed.

### Direction-Only Channels

You can restrict how a function uses a channel:

```go
func emailSender(emailChan <-chan string, done chan<- bool) {
    defer func() {
        done <- true
    }()

    for email := range emailChan {
        fmt.Println("Sending email to:", email)
        time.Sleep(time.Second)
    }
}
```

- `emailChan <-chan string`: **receive-only** inside this function.
- `done chan<- bool`: **send-only** inside this function.

This makes the intent clear and prevents accidental misuse.

### Synchronization with Channels

The `done` channel is used to know when work is finished:

```go
done := make(chan bool)

go emailSender(emailChan, done)

// send values to emailChan
for i := 0; i < 5; i++ {
    emailChan <- fmt.Sprintf("%d@gmail.com", i)
}

fmt.Println("done sending..")
close(emailChan) // important: close so the receiver's range loop can end
<-done            // wait for emailSender to signal completion
```

- `close(emailChan)` tells the receiver that no more values will come.
- `<-done` blocks until `emailSender` sends `true`.

### `select` with Multiple Channels

`select` lets you wait on multiple channel operations and handle whichever is ready first:

```go
chan1 := make(chan int)
chan2 := make(chan string)

go func() {
    chan1 <- 10
}()

go func() {
    chan2 <- "hello"
}()

for i := 0; i < 2; i++ {
    select {
    case chan1Val := <-chan1:
        fmt.Println("Received from chan1:", chan1Val)
    case chan2Val := <-chan2:
        fmt.Println("Received from chan2:", chan2Val)
    }
}
```

- `select` chooses a ready case at runtime.
- Useful when reading from several channels or implementing timeouts.

## Run The Example

Use PowerShell and run:

```powershell
cd "D:\my\GUNJAN\Gunjan Go\learning\channels"
go run main.go
```

You will see output similar to (order/timing may vary):

```text
Sending email to: 0@gmail.com
Sending email to: 1@gmail.com
Sending email to: 2@gmail.com
Sending email to: 3@gmail.com
Sending email to: 4@gmail.com
done sending..
Received from chan2: hello
Received from chan1: 10
```

The order of the last two lines may change because of concurrency.

## Notes and Learning Tips

- Always think about who sends and who receives on a channel; unbalanced sends/receives can cause deadlocks.
- Close channels from the sender side when no more values will be sent (especially when the receiver is using `for range`).
- Use direction-only channel parameters and `select` to write clearer, safer concurrent code.
