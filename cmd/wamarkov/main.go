package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/2zqa/whatsapp-markov-chat/internal/markov"
	"github.com/2zqa/whatsapp-markov-chat/internal/whatsapp"

	"github.com/mb-14/gomarkov"
)

const (
	retryLimit  = 30
	markovOrder = 2
)

func main() {
	// Parse the command line flags
	filepath := flag.String("file", "", "path to the WhatsApp chat export file")
	flag.Parse()

	if *filepath == "" {
		log.Fatal("Please provide the path to the WhatsApp chat export file using the -file flag")
	}

	// Parse the chat file
	messages, err := whatsapp.ParseChat(*filepath)
	if err != nil {
		log.Fatalf("Error parsing chat file: %v", err)
	}

	// Create a Markov chain
	chain := gomarkov.NewChain(markovOrder)
	for _, message := range messages {
		chain.Add(strings.Fields(message.Message))
	}
	tokens := markov.CreateTokenSlice(chain)

	// Main loop to generate messages
	fmt.Print("Press enter to generate a new message...")
	var generatedMessage string
	for {
		fmt.Scanln()
		generatedMessage = generateUniqueMessage(chain, tokens, messages)
		fmt.Print(generatedMessage)
	}
}

func isMessageInList(message string, messages []whatsapp.Message) bool {
	for _, m := range messages {
		if m.Message == message {
			return true
		}
	}
	return false
}

func generateUniqueMessage(chain *gomarkov.Chain, tokens []string, messages []whatsapp.Message) string {
	var generatedMessage string
	for i := 0; i < retryLimit; i++ {
		generatedMessage = markov.Generate(chain, tokens)
		if !isMessageInList(generatedMessage, messages) {
			break
		}
	}
	return generatedMessage
}
