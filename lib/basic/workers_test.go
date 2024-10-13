package basic

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"
	"time"
)

// TestWorker_DoWork tests the DoWork method of the Worker struct.
func TestWorker_DoWork(t *testing.T) {
	var wg sync.WaitGroup
	worker := NewWorker(1)

	wg.Add(1) // Prepare to wait for the worker to finish

	go func() {
		defer wg.Done()    // Decrement the WaitGroup counter when this goroutine completes
		worker.DoWork(&wg) // Call DoWork with the WaitGroup
	}()

	wg.Wait() // Wait for the worker to complete
}

// TestRunWorkers tests the RunWorkers function for expected behavior.
func TestRunWorkers(t *testing.T) {
	var buffer bytes.Buffer
	oldOutput := log.Writer()      // Save the old log writer
	defer log.SetOutput(oldOutput) // Restore the old log writer
	log.SetOutput(&buffer)         // Redirect log output to buffer

	start := time.Now()
	RunWorkers(3) // Run 3 workers
	duration := time.Since(start)

	// Check if the total duration is greater than the sum of worker sleep times
	if duration < 6*time.Second {
		t.Errorf("Expected total duration to be at least 6 seconds, got %v", duration)
	}

	// Check for expected output
	output := buffer.String()
	if !strings.Contains(output, "All workers done") {
		t.Error("Expected output to contain 'All workers done'")
	}
}

// TestRunWorkersWithDifferentCounts tests the RunWorkers function with different worker counts.
func TestRunWorkersWithDifferentCounts(t *testing.T) {
	tests := []struct {
		numWorkers   int
		expectedTime time.Duration
	}{
		{1, 2 * time.Second}, // 1 worker takes approximately 2 seconds
		{2, 2 * time.Second}, // 2 workers still take about 2 seconds
		{3, 2 * time.Second}, // 3 workers still take about 2 seconds
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("numWorkers=%d", tt.numWorkers), func(t *testing.T) {
			start := time.Now()
			RunWorkers(tt.numWorkers)
			duration := time.Since(start)

			if duration < tt.expectedTime {
				t.Errorf("Expected duration to be at least %v, got %v", tt.expectedTime, duration)
			}
		})
	}
}
