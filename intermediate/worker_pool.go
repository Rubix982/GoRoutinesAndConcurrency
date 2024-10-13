package intermediate

import (
	"fmt"
	"sync"
)

// Job represents a single unit of work.
type Job struct {
	ID int
}

// The Result represents the processed result of a job.
type Result struct {
	JobID int
	Value int
	Err   error
}

// WorkerPool processes jobs concurrently using a specified number of workers.
// It accepts a slice of jobs, a number of workers, and a function to process each job.
// It returns the results of the processed jobs or an error if something goes wrong.
func WorkerPool(jobs []Job, workerCount int, processFunc func(Job) (int, error)) ([]Result, error) {
	if workerCount <= 0 {
		return nil, fmt.Errorf("worker count must be greater than zero")
	}
	if len(jobs) == 0 {
		return []Result{}, nil // No jobs to process
	}

	jobChannel := make(chan Job, len(jobs))
	resultChannel := make(chan Result, len(jobs))
	var wg sync.WaitGroup

	// Spawn workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(jobChannel, resultChannel, processFunc, &wg)
	}

	// Add jobs to the channel
	for _, job := range jobs {
		jobChannel <- job
	}
	close(jobChannel) // Close job channel to signal no more jobs

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(resultChannel) // Close results channel when done
	}()

	// Collect results
	var results []Result
	for result := range resultChannel {
		results = append(results, result)
	}

	return results, nil
}

// worker processes jobs from the job channel and sends the processed results to the result channel.
func worker(jobs <-chan Job, results chan<- Result, processFunc func(Job) (int, error), wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the worker is done
	for job := range jobs {
		value, err := processFunc(job)
		results <- Result{JobID: job.ID, Value: value, Err: err}
	}
}
