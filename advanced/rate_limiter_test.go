package advanced

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	rateLimiter := NewRateLimiter(2, time.Second) // 2 actions per second

	allowedActions := 0
	duration := 2 * time.Second // Test for 2 seconds
	start := time.Now()

	// Capture the time allowed actions
	for time.Since(start) < duration {
		rateLimiter.Allow()
		allowedActions++
	}

	// Check if we allowed the expected number of actions
	expectedActions := 4 // 2 actions per second * 2 seconds
	if allowedActions != expectedActions {
		t.Errorf("Expected %d allowed actions, got %d", expectedActions, allowedActions)
	}
}
