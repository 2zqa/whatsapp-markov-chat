package whatsapp

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"time"
)

var messagePattern = regexp.MustCompile(`(\d{2}-\d{2}-\d{4} \d{2}:\d{2}) - ([^:]+): (.*)`)

// Message represents a single message in a WhatsApp chat
type Message struct {
	Timestamp time.Time
	// Name is the name of the person who sent the message
	Name string
	// Message is the content of the message. It may contain newlines.
	Message string
}

func ParseChat(filepath string) ([]Message, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var messages []Message
	scanner := bufio.NewScanner(file)
	var currentMessage *Message

	for scanner.Scan() {
		line := scanner.Text()
		match := messagePattern.FindStringSubmatch(line)
		if match != nil {
			if currentMessage != nil {
				messages = append(messages, *currentMessage)
			}
			timestamp, _ := time.Parse("02-01-2006 15:04", match[1])
			name := strings.TrimSpace(match[2])
			message := strings.TrimSpace(match[3])
			currentMessage = &Message{Timestamp: timestamp, Name: name, Message: message}
		} else if currentMessage != nil {
			currentMessage.Message += "\n" + strings.TrimSpace(line)
		}
	}

	if currentMessage != nil {
		messages = append(messages, *currentMessage)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
