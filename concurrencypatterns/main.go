package main

import (
	"fmt"
	"sync"
	"time"
)


type Result struct{
	Value string	
	Err error
}

// func worker(url string, wg *sync.WaitGroup, resChan chan Result) {
// 	defer wg.Done()
// 	time.Sleep(50 * time.Millisecond)

// 	fmt.Printf("image processed: %s\n", url)
// 	resChan <- Result{
// 		Value: url,
// 		Err:   nil,
// 	}
// }

func worker(jobsChan chan string, wg *sync.WaitGroup, resChan chan Result) {
	defer wg.Done()
	for job := range jobsChan {
		time.Sleep(50 * time.Millisecond)
		// fmt.Println("Image Processed:", job)

		resChan <- Result{
			Value: "Processed " + job,
			Err:   nil,
		}
	}

	fmt.Println("Worker Shutting down")
}

func main() {
	/////// Patterns in Concurrency ///////
	// 1. Fan out / Fan in pattern

	// Fan Out pattern - multiple goroutines doing work concurrently
	// Fan In pattern - collecting results from multiple goroutines
	// var wg sync.WaitGroup
	// resChan := make(chan Result, 5)
	// wg.Add(2)
	// stTime := time.Now()
	// // Fan Out
	// go worker("http://image1.com", &wg, resChan)
	// go worker("http://image2.com", &wg, resChan)
	// wg.Wait()
	// close(resChan)

	// // Fan In
	// for res := range resChan {
	// 	fmt.Println("Result received:", res.Value)
	// }

	// fmt.Println("Time taken:", time.Since(stTime))


	// 2. Worker Pool pattern

	// there is a possibility that 1000s. 100000s of images need to be processed
	// creating that many goroutines will lead to resource exhaustion
	// instead we can create a pool of workers (fixed number of goroutines) to process the jobs

	jobs := []string{
		"http://image1.com",
		"http://image2.com",
		"http://image3.com",
		"http://image4.com",
		"http://image5.com",
		"http://image6.com",
		"http://image7.com",
		"http://image8.com",
		"http://image9.com",
		"http://image10.com",
		"http://image11.com",
		"http://image12.com",
		"http://image13.com",
		"http://image14.com",
		"http://image15.com",
		"http://image16.com",
		"http://image17.com",
		"http://image18.com",
		"http://image19.com",
		"http://image20.com",
	}
	var wg sync.WaitGroup
	resChan := make(chan Result, 50)
	jobsChan := make(chan string, len(jobs))
	totalWorkers := 5

	// for _, job := range jobs {
	// 	wg.Add(1)
	// 	go worker(job, &wg, resChan)
	// }
	for i:=1; i<=totalWorkers; i++{
		wg.Add(1)
		go worker(jobsChan ,&wg, resChan)

	}

	stTime := time.Now()
    
	go func(){
		wg.Wait()
		close(resChan)
	}()

	// sending jobs to jobsChan

	for i:=0; i<len(jobs); i++{
		jobsChan <- jobs[i]
	}
	close(jobsChan)

	for res := range resChan {
		fmt.Println("Job Completed:", res.Value)
	}

	fmt.Println("Time taken:", time.Since(stTime))

	
}
