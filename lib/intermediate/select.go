package intermediate

import (
	"fmt"
	"time"
)

// Message represents a message received from a channel.
type Message struct {
	Content string
	Source  string
}

// Select demonstrating how to receive messages from multiple channels using the select statement.
func Select() {
	// Create channels for messages
	ch1 := make(chan Message)
	ch2 := make(chan Message)

	// Launch goroutines to send messages to channels after a delay
	go SendMessage(ch1, "Message from ch1", 1*time.Second)
	go SendMessage(ch2, "Message from ch2", 2*time.Second)

	// Receive messages from both channels
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			HandleMessage(msg)
		case msg := <-ch2:
			HandleMessage(msg)
		}
	}
}

// SendMessage simulates sending a message to a channel after a specified delay.
func SendMessage(ch chan Message, content string, delay time.Duration) {
	time.Sleep(delay)
	ch <- Message{Content: content, Source: "Channel"}
}

// HandleMessage processes the received message and prints it.
func HandleMessage(msg Message) {
	fmt.Println("Received:", msg.Content)
}

// SendMessages sends multiple messages to a specified channel with varying delays.
// This allows for more flexibility and testing with different message sources.
func SendMessages(ch chan Message, messages []string, delays []time.Duration) {
	for i, content := range messages {
		go SendMessage(ch, content, delays[i])
	}
}
