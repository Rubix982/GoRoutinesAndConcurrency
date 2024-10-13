package intermediate

import (
	"bytes"
	"io"
	"os"
)

// CaptureOutput - Helper function to capture output
func CaptureOutput(f func()) string {
	// Create a pipe
	r, w, _ := os.Pipe()

	// Save the original stdout
	originalStdout := os.Stdout
	defer func() {
		os.Stdout = originalStdout // Restore original stdout
	}()

	// Set stdout to the pipe
	os.Stdout = w

	// Run the function to capture output
	f()

	// Close the writer
	if err := w.Close(); err != nil {
		return ""
	}

	// Read from the pipe
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		return ""
	} // Copy the output from the pipe to the buffer

	return buf.String() // Return captured output
}
