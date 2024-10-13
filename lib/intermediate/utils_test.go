package intermediate

import (
	"fmt"
	"testing"
)

// TestCaptureOutput verifies that CaptureOutput correctly captures output.
func TestCaptureOutput(t *testing.T) {
	expected := "Hello, world!\n" // Expecting this output with a newline

	// Define a function that prints to stdout
	testFunc := func() {
		fmt.Println("Hello, world!") // Use fmt.Println instead of println
	}

	// Capture the output of the test function
	output := CaptureOutput(testFunc)

	// Verify that the captured output matches the expected output
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}
