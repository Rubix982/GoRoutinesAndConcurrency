package basic

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// TestExampleGoroutines tests the ExampleGoroutines function for concurrent message printing.
func TestExampleGoroutines(t *testing.T) {
	// Create a pipe to capture output
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	// Save original stdout
	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }() // Restore original stdout

	// Redirect stdout to pipe
	os.Stdout = w

	// Test with various messages
	messages := []string{"Goroutine 1", "Goroutine 2", "Goroutine 3"}
	ExampleGoroutines(messages...)

	// Close the writer to flush the pipe
	w.Close()

	// Read the output from the reader
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatal(err)
	}

	// Check if all expected messages are present in the output
	outputString := buf.String()
	for _, msg := range messages {
		if !strings.Contains(outputString, msg) {
			t.Errorf("Expected to find '%s' in output, got:\n%s", msg, outputString)
		}
	}
}

// TestRunGoroutines tests the RunGoroutines function for proper message execution.
func TestRunGoroutines(t *testing.T) {
	// Create a pipe to capture output
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	// Save original stdout
	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }() // Restore original stdout

	// Redirect stdout to pipe
	os.Stdout = w

	RunGoroutines()

	// Close the writer to flush the pipe
	w.Close()

	// Read the output from the reader
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatal(err)
	}

	// Check if all expected messages are present in the output
	expectedMessages := []string{"Goroutine 1", "Goroutine 2", "Goroutine 3"}
	outputString := buf.String()
	for _, msg := range expectedMessages {
		if !strings.Contains(outputString, msg) {
			t.Errorf("Expected to find '%s' in output, got:\n%s", msg, outputString)
		}
	}
}

// TestPrintMessage tests the PrintMessage function to verify its output.
func TestPrintMessage(t *testing.T) {
	// Create a pipe to capture output
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	// Save original stdout
	originalStdout := os.Stdout
	defer func() { os.Stdout = originalStdout }() // Restore original stdout

	// Redirect stdout to pipe
	os.Stdout = w

	// Test with a single message
	message := "Hello, World!"
	PrintMessage(message)

	// Close the writer to flush the pipe
	w.Close()

	// Read the output from the reader
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatal(err)
	}

	// Check if the output is correct
	outputString := buf.String()
	if outputString != message+"\n" {
		t.Errorf("Expected output '%s', got '%s'", message, outputString)
	}
}
