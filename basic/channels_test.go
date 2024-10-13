package basic

import (
	"testing"
)

// TestSendMessage tests sending a single message.
func TestSendMessage(t *testing.T) {
	sender := NewMessageSender(1)
	defer sender.Wait() // Ensure resources are cleaned up

	// Test sending a simple message
	msg := "Hello from Test"
	result, err := sender.SendMessage(msg)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Content != msg {
		t.Errorf("Expected message content '%s', got '%s'", msg, result.Content)
	}
}

// TestMultipleMessages tests sending multiple messages.
func TestMultipleMessages(t *testing.T) {
	sender := NewMessageSender(3)
	defer sender.Wait() // Ensure resources are cleaned up

	expectedMessages := []string{
		"First Message",
		"Second Message",
		"Third Message",
	}

	for _, expected := range expectedMessages {
		result, err := sender.SendMessage(expected)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.Content != expected {
			t.Errorf("Expected message content '%s', got '%s'", expected, result.Content)
		}
	}
}

// TestEmptyMessage tests sending an empty message.
func TestEmptyMessage(t *testing.T) {
	sender := NewMessageSender(1)
	defer sender.Wait() // Ensure resources are cleaned up

	result, err := sender.SendMessage("")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.Content != "" {
		t.Errorf("Expected empty message, got '%s'", result.Content)
	}
}

// TestBufferedChannel tests the behavior of buffered channels.
func TestBufferedChannel(t *testing.T) {
	sender := NewMessageSender(2)
	defer sender.Wait() // Ensure resources are cleaned up

	messages := []string{
		"Message 1",
		"Message 2",
		"Message 3",
	}

	for _, msg := range messages {
		result, err := sender.SendMessage(msg)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result.Content != msg {
			t.Errorf("Expected message content '%s', got '%s'", msg, result.Content)
		}
	}
}

// TestChannelClose checks that the channel is properly closed after all messages are sent.
func TestChannelClose(t *testing.T) {
	sender := NewMessageSender(1)

	// Start sending a message and close the sender
	message, err := sender.SendMessage("Closing message")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if message.Content != "Closing message" {
		t.Errorf("Expected message 'Closing message', got '%s'", message.Content)
	}

	// Wait for all goroutines to finish
	sender.Wait()

	// Attempting to send after wait should cause a panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when sending after Wait, but got none")
		}
	}()

	// Attempting to send a message after the Wait should cause a panic
	_, err = sender.SendMessage("Should fail") // This line should cause a panic
	// Note: We don't need to check the error here since we expect a panic
}
