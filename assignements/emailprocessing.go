// Requirements (What you must build)
// 1. A dispatcher goroutine
// Reads incoming email jobs from a main job channel
// Assigns jobs to a pool of worker goroutines

// 2. A worker pool (N workers)
// Each worker:
// Receives jobs via its own worker channel
// Simulates sending email (sleep)
// Prints: "Worker X sent email to <address>"

// 3. Use WaitGroup to wait for ALL workers to finish
// 4. Use buffered channels for job queue
// 5. Support graceful shutdown

// When main closes the Job channel:
// Dispatcher must not panic
// Workers must stop after finishing their current tasks
// Main waits for all workers to exit
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type EmailJob struct {
// 	ID    int
// 	Email string
// }

// func worker(id int, jobs <-chan EmailJob, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for job := range jobs {
// 		email := job.Email
// 		fmt.Println("Worker", id, "sending email to:", email)
// 		time.Sleep(time.Second)
// 	}

// 	fmt.Println("Worker", id, "has no more jobs, exiting.")
// }

// func dispatcher(jobq <-chan EmailJob, nWorkers int, wgg *sync.WaitGroup) {
// 	defer wgg.Done()

// 	var wg sync.WaitGroup
// 	workerChans := make([]chan EmailJob, nWorkers)
// 	for i := 0; i < nWorkers; i++ {
// 		workerChans[i] = make(chan EmailJob)
// 		wg.Add(1)
// 		go worker(i, workerChans[i], &wg) // start worker goroutine
// 	}
// 	for job := range jobq {
// 		wId := job.ID % nWorkers
// 		workerChans[wId] <- job
// 	}
// 	for i := 0; i < nWorkers; i++ {
// 		close(workerChans[i])
// 	}
// 	wg.Wait()
// }

// func main() {

// 	jobQ := make(chan EmailJob, 20) // buffered channel
// 	nWorkers := 3

// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go dispatcher(jobQ, nWorkers, &wg)

// 	for i := 0; i < 10; i++ { // sending 10 email jobs
// 		jobQ <- EmailJob{
// 			ID:    i,
// 			Email: fmt.Sprintf("user%d@example.com", i),
// 		}
// 	}

// 	close(jobQ)
// 	wg.Wait()

// }
package main
import (
	"fmt"
	"sync"
	"time"
)

type EmailJob struct {
	ID    int
	Email string
}

// worker reads from the shared jobQueue and obeys global rate limiting
func worker(id int, jobQueue <-chan EmailJob, wg *sync.WaitGroup, limiter <-chan time.Time) {
	defer wg.Done()
	for job := range jobQueue {
		<-limiter // global rate limit before processing

		fmt.Println("Worker", id, "sending email to:", job.Email)
		time.Sleep(time.Second) // simulate slow email API
	}
	fmt.Println("Worker", id, "exiting...")
}

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
