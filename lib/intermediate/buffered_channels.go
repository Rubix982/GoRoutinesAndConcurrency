package intermediate

import (
	"fmt"
	"sync"
)

// Counter safely increments a shared integer value.
type Counter struct {
	mu    sync.Mutex // Mutex to protect shared data
	value int        // The counter-value
}

// NewCounter initializes and returns a new Counter instance.
func NewCounter() *Counter {
	return &Counter{}
}

// Increment increases the counter by 1.
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() // Ensure the mutex is unlocked even if an error occurs
	c.value++
}

// Value returns the current counter value safely.
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock() // Ensure the mutex is unlocked
	return c.value
}

// Display prints the current counter-value in a formatted way.
// This function can be useful for debugging or logging.
func (c *Counter) Display() {
	fmt.Printf("Current Counter Value: %d\n", c.Value())
}

// Reset sets the counter-value back to zero.
// This method provides a way to restart counting.
func (c *Counter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = 0
}

// SetValue sets the counter to a specific value.
// This method allows more flexibility in modifying the counter.
func (c *Counter) SetValue(newValue int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = newValue
}

// GetValue safely retrieves the current value.
// This method provides a clear interface for accessing the counter.
func (c *Counter) GetValue() int {
	return c.Value()
}

// GoRoutineIncrement increments the counter using multiple goroutines.
// This method showcases the concurrent capabilities of the Counter.
func (c *Counter) GoRoutineIncrement(n int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure the WaitGroup counter is decremented
	for i := 0; i < n; i++ {
		c.Increment()
	}
}
