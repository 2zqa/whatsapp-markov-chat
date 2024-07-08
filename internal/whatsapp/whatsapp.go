package whatsapp

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"time"
)

const whatsappTimestampLayout = "02-01-2006 15:04"

var messagePattern = regexp.MustCompile(`(\d{2}-\d{2}-\d{4} \d{2}:\d{2}) - ([^:]+): (.*)`)

// Message represents a single message in a WhatsApp chat
type Message struct {
	// Timestamp is the moment the message was sent
	Timestamp time.Time
	// Name is the name of the person who sent the message
	Name string
	// Message is the content of the message. May contain newlines.
	Message string
}

func ParseChat(filepath string) ([]Message, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var messages []Message
	var currentMessage *Message
	for scanner.Scan() {
		line := scanner.Text()
		match := messagePattern.FindStringSubmatch(line)
		if match != nil {
			if currentMessage != nil {
				messages = append(messages, *currentMessage)
			}
			timestamp, _ := time.Parse(whatsappTimestampLayout, match[1])
			name := match[2]
			message := match[3]
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
