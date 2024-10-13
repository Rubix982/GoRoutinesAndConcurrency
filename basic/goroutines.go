package basic

import (
	"fmt"
	"sync"
)

// ExampleGoroutines demonstrates how to run multiple goroutines concurrently.
// It accepts a variable number of messages to print, showcasing concurrent execution.
// Each message is printed by a separate goroutine.
func ExampleGoroutines(messages ...string) {
	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish

	for _, message := range messages {
		wg.Add(1) // Increment the WaitGroup counter

		go func(msg string) {
			defer wg.Done()   // Decrement the counter when the goroutine completes
			PrintMessage(msg) // Call a separate function to handle printing
		}(message) // Pass the message to avoid closure issues
	}

	wg.Wait() // Wait for all goroutines to complete
}

// PrintMessage is a helper function that prints a message to standard output.
// It can be mocked in tests to verify the output without printing to the console.
func PrintMessage(msg string) {
	fmt.Println(msg)
}

// RunGoroutines is a convenience function to demonstrate the execution of ExampleGoroutines.
func RunGoroutines() {
	messages := []string{"Goroutine 1", "Goroutine 2", "Goroutine 3"}
	ExampleGoroutines(messages...) // Pass messages to ExampleGoroutines
}
