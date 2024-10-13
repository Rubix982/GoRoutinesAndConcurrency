package intermediate

import (
	"fmt"
	"testing"
)

// mockProcessFunc is a mock function for testing job processing.
func mockProcessFunc(job Job) (int, error) {
	if job.ID < 0 {
		return 0, fmt.Errorf("invalid job ID")
	}
	return job.ID * 2, nil // Example processing: double the job ID
}

// TestWorkerPool tests the WorkerPool function with valid input.
func TestWorkerPool(t *testing.T) {
	jobs := []Job{{ID: 1}, {ID: 2}, {ID: 3}}
	workerCount := 2

	results, err := WorkerPool(jobs, workerCount, mockProcessFunc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check results
	expectedResults := []Result{
		{JobID: 1, Value: 2, Err: nil},
		{JobID: 2, Value: 4, Err: nil},
		{JobID: 3, Value: 6, Err: nil},
	}

	if len(results) != len(expectedResults) {
		t.Errorf("Expected %d results, got %d", len(expectedResults), len(results))
	}

	for i, result := range results {
		if result != expectedResults[i] {
			t.Errorf("Expected result %v, got %v", expectedResults[i], result)
		}
	}
}

// TestWorkerPoolWithInvalidJobs tests the WorkerPool with an invalid job ID.
func TestWorkerPoolWithInvalidJobs(t *testing.T) {
	jobs := []Job{{ID: -1}} // Invalid job
	workerCount := 1

	results, err := WorkerPool(jobs, workerCount, mockProcessFunc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check results
	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}
	if results[0].Err == nil {
		t.Errorf("Expected error for invalid job ID, got result %v", results[0])
	} else if results[0].JobID != -1 || results[0].Value != 0 {
		t.Errorf("Expected result for job ID -1, got result %v", results[0])
	}
}

// TestWorkerPoolWithZeroWorkers tests the WorkerPool with zero workers.
func TestWorkerPoolWithZeroWorkers(t *testing.T) {
	jobs := []Job{{ID: 1}, {ID: 2}}
	workerCount := 0

	_, err := WorkerPool(jobs, workerCount, mockProcessFunc)
	if err == nil {
		t.Fatalf("Expected error for zero workers, got none")
	}
}

// TestWorkerPoolWithNegativeJobs tests the WorkerPool with a slice containing negative job IDs.
func TestWorkerPoolWithNegativeJobs(t *testing.T) {
	jobs := []Job{{ID: -1}, {ID: -2}}
	workerCount := 2

	results, err := WorkerPool(jobs, workerCount, mockProcessFunc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	for _, result := range results {
		if result.Err == nil {
			t.Errorf("Expected an error for invalid job ID, got result %v", result)
		}
	}
}

// TestWorkerPoolWithNoJobs tests the WorkerPool with an empty job list.
func TestWorkerPoolWithNoJobs(t *testing.T) {
	var jobs []Job
	workerCount := 2

	results, err := WorkerPool(jobs, workerCount, mockProcessFunc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(results) != 0 {
		t.Errorf("Expected no results for empty job list, got %d results", len(results))
	}
}
