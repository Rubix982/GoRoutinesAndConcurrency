package advanced

import (
	"fmt"
	"sync"
	"time"
)

// BoundedSemaphore controls the number of concurrent goroutines using a semaphore
func BoundedSemaphore(tasks []int, maxConcurrency int) {
	sem := make(chan struct{}, maxConcurrency) // Bounded semaphore
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		go func(task int) {
			defer wg.Done()

			sem <- struct{}{}        // Acquire semaphore slot
			defer func() { <-sem }() // Release semaphore slot after task completes

			fmt.Printf("Processing task %d\n", task)
			time.Sleep(1 * time.Second) // Simulate some processing time
		}(task)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All tasks completed")
}
