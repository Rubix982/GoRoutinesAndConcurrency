package advanced

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestBoundedSemaphore(t *testing.T) {
	// Create a pipe to capture stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	old := os.Stdout
	defer func() { os.Stdout = old }()
	os.Stdout = w // Redirect stdout to the writer

	// Define test tasks
	tasks := []int{1, 2, 3, 4, 5}
	maxConcurrency := 2

	// Use a wait group to wait for the goroutine to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Run the bounded semaphore function in a goroutine
	go func() {
		defer wg.Done() // Signal that this goroutine is done
		BoundedSemaphore(tasks, maxConcurrency)
	}()

	// Close the writer to flush the output after the goroutine completes
	go func() {
		wg.Wait() // Wait for the bounded semaphore function to finish
		w.Close() // Close the writer to flush output
	}()

	// Read the output from the reader
	var output strings.Builder
	if _, err := io.Copy(&output, r); err != nil {
		t.Fatalf("Failed to read output: %v", err)
	}

	// Check output for expected processing
	lines := strings.Split(output.String(), "\n")
	if len(lines) < len(tasks) {
		t.Errorf("Expected at least %d lines of output, got %d", len(tasks), len(lines))
	}

	// Check that the correct number of tasks are processed
	taskCount := 0
	for _, line := range lines {
		if strings.Contains(line, "Processing task") {
			taskCount++
		}
	}

	// Validate that all tasks were processed
	if taskCount != len(tasks) {
		t.Errorf("Expected to process %d tasks, but processed %d", len(tasks), taskCount)
	}
}
