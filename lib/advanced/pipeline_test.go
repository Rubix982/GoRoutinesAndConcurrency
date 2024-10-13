package advanced

import (
	"testing"
)

func TestPipeline(t *testing.T) {
	// Create an input channel with test data
	input := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)
	}()

	// Set up the pipeline stages
	stage1 := MultiplyByTwo(input)
	stage2 := AddTen(stage1)

	// Collect results
	expectedResults := []int{12, 14, 16, 18, 20} // (1*2 + 10), (2*2 + 10), ..., (5*2 + 10)
	var results []int

	for result := range stage2 {
		results = append(results, result)
	}

	// Check if results match the expected results
	if len(results) != len(expectedResults) {
		t.Errorf("Expected %d results, got %d", len(expectedResults), len(results))
	}

	for i, expected := range expectedResults {
		if results[i] != expected {
			t.Errorf("At index %d: expected %d, got %d", i, expected, results[i])
		}
	}
}
