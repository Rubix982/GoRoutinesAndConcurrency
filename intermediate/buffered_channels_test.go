package intermediate

import (
	"sync"
	"testing"
)

// TestCounter_IncrementAndValue tests the Increment and Value methods.
func TestCounter_IncrementAndValue(t *testing.T) {
	counter := NewCounter()

	// Increment counter in a single goroutine
	counter.Increment()
	if counter.Value() != 1 {
		t.Errorf("Expected counter to be 1, got %d", counter.Value())
	}

	// Increment counter in multiple goroutines
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()

	if counter.Value() != 1001 {
		t.Errorf("Expected counter to be 1001, got %d", counter.Value())
	}
}

// TestCounter_Display tests the Display method.
func TestCounter_Display(t *testing.T) {
	counter := NewCounter()
	counter.Increment() // Increment to have a value

	// Redirect stdout for testing purposes
	output := CaptureOutput(func() {
		counter.Display()
	})

	expectedOutput := "Current Counter Value: 1\n" // After one increment
	if output != expectedOutput {
		t.Errorf("Expected output to be '%s', got '%s'", expectedOutput, output)
	}
}

// TestCounter_Reset tests the Reset method.
func TestCounter_Reset(t *testing.T) {
	counter := NewCounter()
	counter.Increment() // Increment to have a value

	// Ensure the counter-value is 1 before reset
	if counter.Value() != 1 {
		t.Errorf("Expected counter to be 1 before reset, got %d", counter.Value())
	}

	counter.Reset() // Reset the counter

	// Ensure the counter-value is now 0
	if counter.Value() != 0 {
		t.Errorf("Expected counter to be 0 after reset, got %d", counter.Value())
	}
}
