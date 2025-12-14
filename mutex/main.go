package main

import (
	"fmt"
	"sync"
)

// mutexes are used to provide safe access to shared resources between goroutines
// they provide locking mechanism to ensure only one goroutine can access the resource at a time

// race condition occurs when multiple goroutines access shared resource concurrently
// and at least one of the accesses is a write operation
// mutex here can be used to prevent race conditions by locking the resource during access



type post struct{
	views int
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup, i int){
	p.mu.Lock()
	p.views++
	fmt.Println("Incremented views to:", p.views)
	// trace go routine no
	fmt.Println("Go routine no: ", i)
	defer func(){
		wg.Done()
		p.mu.Unlock()

	}()
}

func main(){
	var wg sync.WaitGroup
    myPost := post{views:0}
    
	for i:= range 10{
		wg.Add(1)
		go myPost.inc(&wg, i)
	}

	wg.Wait()

	fmt.Println("Views:", myPost.views)
}