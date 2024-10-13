package intermediate

import (
	"testing"
	"time"
)

func TestContextExample_Completion(t *testing.T) {
	// Set up a context with a timeout longer than the goroutine's sleep time
	ctx, cancel := NewContextWithTimeout(2 * time.Second)
	defer cancel()

	// Capture the output
	output := CaptureOutput(func() {
		ContextExample(ctx)
	})

	expected := "Work completed\n" // Expect the output to have a newline
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func TestContextExample_Cancellation(t *testing.T) {
	// Set up a context with a timeout shorter than the goroutine's sleep time
	ctx, cancel := NewContextWithTimeout(500 * time.Millisecond)
	defer cancel()

	// Capture the output
	output := CaptureOutput(func() {
		ContextExample(ctx)
	})

	expected := "Operation cancelled: context deadline exceeded\n" // Expect the output to have a newline
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}
