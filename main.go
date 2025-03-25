package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: Job channel closed. Exiting.\n", id)
				return
			}
			// Simulate variable processing time
			delay := time.Duration(rand.Intn(500)+200) * time.Millisecond
			time.Sleep(delay)

			fmt.Printf("Worker %d: Processed job %d in %v\n", id, job.ID, delay)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const (
		numJobs    = 10
		numWorkers = 3
	)

	jobs := make(chan Job)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Dispatcher: send jobs
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- Job{ID: j}
			fmt.Printf("Dispatcher: Sent job %d\n", j)
			time.Sleep(150 * time.Millisecond) // Simulate job arrival rate
		}
		close(jobs)
	}()

	// Wait for all workers to finish
	wg.Wait()
	fmt.Println("All jobs processed.")
}
