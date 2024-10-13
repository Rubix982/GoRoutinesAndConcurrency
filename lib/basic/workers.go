package basic

import (
	"fmt"
	"sync"
	"time"
)

// Worker represents a worker that performs a task.
type Worker struct {
	ID int
}

// NewWorker initializes a new Worker instance.
func NewWorker(id int) *Worker {
	return &Worker{ID: id}
}

// DoWork simulates the work done by the worker and accepts a WaitGroup for completion signaling.
func (w *Worker) DoWork(wg *sync.WaitGroup) {
	defer wg.Done() // Notify that this worker has completed
	fmt.Printf("Worker %d started\n", w.ID)
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Printf("Worker %d finished\n", w.ID)
}

// RunWorkers launches multiple workers concurrently and waits for them to complete.
func RunWorkers(numWorkers int) {
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Increment WaitGroup counter
		worker := NewWorker(i)
		go worker.DoWork(&wg) // Launch worker as Goroutine
	}

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All workers done")
}
