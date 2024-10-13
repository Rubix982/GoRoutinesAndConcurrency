package basic

import (
	"fmt"
	"sync"
)

// SafeCounter provides a thread-safe counter using a mutex.
type SafeCounter struct {
	mu      sync.Mutex // Mutex to protect the shared counter
	counter int        // The shared counter
}

// NewSafeCounter initializes and returns a new SafeCounter.
func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

// Increment safely increments the counter by 1.
func (sc *SafeCounter) Increment() {
	sc.mu.Lock()         // Acquire the mutex to protect shared data
	defer sc.mu.Unlock() // Ensure the mutex is released after the function completes

	sc.counter++            // Increment the counter
	fmt.Println(sc.counter) // Logging for visibility (can be removed in production)
}

// Get returns the current value of the counter safely.
func (sc *SafeCounter) Get() int {
	sc.mu.Lock()         // Acquire the mutex to protect shared data
	defer sc.mu.Unlock() // Ensure the mutex is released after the function completes

	return sc.counter // Return the current counter-value
}

// Reset resets the counter to zero safely.
func (sc *SafeCounter) Reset() {
	sc.mu.Lock()         // Acquire the mutex to protect shared data
	defer sc.mu.Unlock() // Ensure the mutex is released after the function completes

	sc.counter = 0 // Reset the counter to zero
}

// Set allows setting the counter to a specific value safely.
func (sc *SafeCounter) Set(value int) {
	sc.mu.Lock()         // Acquire the mutex to protect shared data
	defer sc.mu.Unlock() // Ensure the mutex is released after the function completes

	sc.counter = value // Set the counter to the provided value
}
