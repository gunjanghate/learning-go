package main

import (
	"fmt"
	"sync"
	// "time"
)

// light weight threads managed by go runtime
// used when you want to perform multiple tasks concurrently
// goroutines are cheaper than OS threads
// goroutines are multiplexed onto OS threads by go runtime scheduler
// when a goroutine is blocked, the scheduler will schedule another goroutine onto the same OS thread
// goroutines communicate using channels


// waitgroups are used to wait for a collection of goroutines to finish
// a waitgroup has a counter that is incremented when a goroutine is started and decremented when a goroutine is finished
// the main goroutine can wait for all goroutines to finish by calling Wait on the waitgroup

func tasks(id int, w *sync.WaitGroup) {
	defer w.Done() // defer: ensure that Done is called when the function exits
	fmt.Println("Task", id, "is starting")
}

func main() {
	// for i := 0; i < 10; i++ {
	// 	go tasks(i) // launch a goroutine for each task
	// 	// fmt.Println("Launched goroutine for task", i)
	// 	// go func(i int) {
	// 	// 	fmt.Println("Task", i, "is starting")
	// 	// }(i)
	// }

	// time.Sleep(time.Second * 2) // why? to allow goroutines to complete before main exits

	/// waitgroups to wait for goroutines to complete

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)        // increment the waitgroup counter
		go tasks(i, &wg) // launch a goroutine for each task
	}


	wg.Wait() // wait for all goroutines to complete
}

// output

// Task 3 is starting
// Task 2 is starting
// Task 0 is starting
// Task 1 is starting







