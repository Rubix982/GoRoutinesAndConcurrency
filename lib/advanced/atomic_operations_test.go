package advanced

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomicOperations(t *testing.T) {
	var counter int64
	var wg sync.WaitGroup

	// Launch 1000 goroutines to increment the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Verify the final value of the counter
	finalCounter := atomic.LoadInt64(&counter)
	if finalCounter != 1000 {
		t.Errorf("Expected counter to be 1000, got %d", finalCounter)
	}
}
