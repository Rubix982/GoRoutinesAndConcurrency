package intermediate

import (
	"testing"
	"time"
)

// TestSelect verifies that the select statement correctly receives messages from multiple channels.
func TestSelect(t *testing.T) {
	ch1 := make(chan Message)
	ch2 := make(chan Message)

	// Mock goroutines to send messages
	go SendMessage(ch1, "Test message from ch1", 1*time.Second)
	go SendMessage(ch2, "Test message from ch2", 2*time.Second)

	// Use a timeout to avoid hanging tests
	timeout := time.After(3 * time.Second)
	msgCount := 0

	for msgCount < 2 {
		select {
		case msg := <-ch1:
			if msg.Content != "Test message from ch1" {
				t.Errorf("Expected 'Test message from ch1', got '%s'", msg.Content)
			}
			msgCount++
		case msg := <-ch2:
			if msg.Content != "Test message from ch2" {
				t.Errorf("Expected 'Test message from ch2', got '%s'", msg.Content)
			}
			msgCount++
		case <-timeout:
			t.Error("Test timed out waiting for messages")
			return
		}
	}
}

// TestSendMessages tests sending multiple messages to a channel with varying delays.
func TestSendMessages(t *testing.T) {
	ch := make(chan Message)
	messages := []string{"Msg1", "Msg2", "Msg3"}
	delays := []time.Duration{1 * time.Second, 2 * time.Second, 3 * time.Second}

	go SendMessages(ch, messages, delays)

	// Collect messages sent
	timeout := time.After(5 * time.Second)
	msgCount := 0

	for msgCount < len(messages) {
		select {
		case msg := <-ch:
			if msgCount < len(messages) && msg.Content != messages[msgCount] {
				t.Errorf("Expected '%s', got '%s'", messages[msgCount], msg.Content)
			}
			msgCount++
		case <-timeout:
			t.Error("Test timed out waiting for messages")
			return
		}
	}
}
