package basic

import (
	"sync"
	"testing"
)

// TestSafeCounter_Increment tests the Increment method of SafeCounter.
func TestSafeCounter_Increment(t *testing.T) {
	counter := NewSafeCounter()
	var wg sync.WaitGroup

	const numGoroutines = 1000

	// Launch multiple goroutines to increment the counter
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait() // Wait for all goroutines to complete
	if finalValue := counter.Get(); finalValue != numGoroutines {
		t.Errorf("Expected counter to be %d, got %d", numGoroutines, finalValue)
	}
}

// TestSafeCounter_Get tests the Get method of SafeCounter.
func TestSafeCounter_Get(t *testing.T) {
	counter := NewSafeCounter()
	counter.Increment() // Increment once

	if val := counter.Get(); val != 1 {
		t.Errorf("Expected counter to be 1, got %d", val)
	}

	// Increment again
	counter.Increment()
	if val := counter.Get(); val != 2 {
		t.Errorf("Expected counter to be 2, got %d", val)
	}
}

// TestSafeCounter_Reset tests the Reset method of SafeCounter.
func TestSafeCounter_Reset(t *testing.T) {
	counter := NewSafeCounter()
	counter.Increment() // Increment once

	counter.Reset() // Reset the counter

	if val := counter.Get(); val != 0 {
		t.Errorf("Expected counter to be 0 after reset, got %d", val)
	}
}

// TestSafeCounter_ConcurrentAccess tests concurrent access and resetting the counter.
func TestSafeCounter_ConcurrentAccess(t *testing.T) {
	counter := NewSafeCounter()
	var wg sync.WaitGroup

	const numGoroutines = 100

	// Launch goroutines to increment the counter
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait() // Wait for all increments to complete

	// Reset the counter and check its value
	counter.Reset()
	if val := counter.Get(); val != 0 {
		t.Errorf("Expected counter to be 0 after reset, got %d", val)
	}
}

// TestSafeCounter_Set tests the Set method of SafeCounter.
func TestSafeCounter_Set(t *testing.T) {
	counter := NewSafeCounter()
	counter.Set(42) // Set the counter to 42

	if val := counter.Get(); val != 42 {
		t.Errorf("Expected counter to be 42, got %d", val)
	}
}

// TestSafeCounter_SetConcurrent tests setting the counter concurrently.
func TestSafeCounter_SetConcurrent(t *testing.T) {
	counter := NewSafeCounter()
	var wg sync.WaitGroup

	const numGoroutines = 10

	// Launch multiple goroutines to set the counter
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			counter.Set(val)
		}(i)
	}

	wg.Wait() // Wait for all goroutines to complete

	// Check that the final value is one of the set values
	if finalValue := counter.Get(); finalValue < 0 || finalValue >= numGoroutines {
		t.Errorf("Expected counter to be one of the set values, got %d", finalValue)
	}
}
