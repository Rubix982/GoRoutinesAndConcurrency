package intermediate

import (
	"context"
	"fmt"
	"time"
)

// ContextExample demonstrates how to use context for cancellation and timeout.
func ContextExample(ctx context.Context) {
	ch := make(chan string)

	// Launch a goroutine that simulates work
	go func() {
		// Simulate work with sleep
		time.Sleep(1 * time.Second)
		ch <- "Work completed"
	}()

	// Use a select statement to wait for either completion or cancellation
	select {
	case result := <-ch:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("Operation cancelled:", ctx.Err())
	}
}

// NewContextWithTimeout creates a new context with a specified timeout duration.
func NewContextWithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}
