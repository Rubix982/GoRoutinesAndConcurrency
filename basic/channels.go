package basic

import (
	"fmt"
	"sync"
)

// ChannelMessage represents a message to be sent over a channel.
type ChannelMessage struct {
	Content string
}

// MessageSender is responsible for sending messages through a channel.
type MessageSender struct {
	ch     chan ChannelMessage
	wg     sync.WaitGroup
	closed bool // New field to indicate if the sender is closed
}

// NewMessageSender initializes a MessageSender instance with a buffered channel.
func NewMessageSender(bufferSize int) *MessageSender {
	return &MessageSender{
		ch: make(chan ChannelMessage, bufferSize),
	}
}

// SendMessage starts a goroutine to send a message and returns the received message.
// It blocks until the message is received to ensure synchronous communication.
func (ms *MessageSender) SendMessage(msgContent string) (ChannelMessage, error) {
	if ms.closed { // Check if the sender is closed
		panic("attempt to send on a closed MessageSender")
	}

	ms.wg.Add(1) // Increment wait group counter

	// Start a goroutine to send the message
	go func(content string) {
		defer ms.wg.Done() // Decrement counter when goroutine completes
		ms.ch <- ChannelMessage{Content: content}
	}(msgContent)

	return <-ms.ch, nil // Block until the message is received from the channel
}

// Wait closes the channel after all messages are sent.
// It is essential to call Wait before exiting to ensure all messages are processed.
func (ms *MessageSender) Wait() {
	ms.wg.Wait()    // Wait for all goroutines to finish
	if !ms.closed { // Check if the channel is already closed
		ms.closed = true // Set closed flag to true
		close(ms.ch)     // Close the channel to signal completion
	}
}

// ExampleChannels demonstrates basic usage of the MessageSender.
func ExampleChannels() {
	sender := NewMessageSender(1) // Create a MessageSender with a buffer size of 1

	// Send a message and print the content
	message, _ := sender.SendMessage("Hello from Goroutine")
	fmt.Println(message.Content)

	sender.Wait() // Ensure all goroutines have finished before exiting
}
