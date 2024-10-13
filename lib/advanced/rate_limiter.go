package advanced

import (
	"fmt"
	"time"
)

// RateLimiter limits the rate of allowed actions over a specified duration.
type RateLimiter struct {
	ticker *time.Ticker
}

// NewRateLimiter creates a new RateLimiter that allows a specified number of actions per interval.
func NewRateLimiter(rate int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		ticker: time.NewTicker(interval / time.Duration(rate)),
	}
}

// Allow waits for the next tick and returns true to indicate an action is allowed.
func (r *RateLimiter) Allow() {
	<-r.ticker.C
}

// Stop stops the rate limiter's ticker.
func (r *RateLimiter) Stop() {
	r.ticker.Stop()
}

// RateLimiterDemo demonstrates usage of the rate limiter.
func RateLimiterDemo() {
	rateLimiter := NewRateLimiter(2, time.Second) // Allow 2 actions per second
	for i := 0; i < 5; i++ {
		rateLimiter.Allow()
		fmt.Println("Allowed at", time.Now())
	}
}
