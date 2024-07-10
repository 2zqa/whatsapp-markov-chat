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

	// Create a Markov chain for every author
	chains := createChains(messages)
	tokens := markov.CreateTokenSlice(markovOrder)

	// Main loop to generate messages
	fmt.Println("Press enter to generate a new message...")
	for {
		fmt.Scanln() // Wait until the user presses enter
		sendMessageForAllParticipants(chains, tokens, messages)
	}
}

func sendMessageForAllParticipants(chains map[string]*gomarkov.Chain, tokens []string, messages []whatsapp.Message) {
	for author, chain := range chains {
		generatedMessage := generateUniqueMessage(chain, tokens, messages)
		fmt.Print(author + ": " + generatedMessage + "\n")
	}
}

func createChains(messages []whatsapp.Message) map[string]*gomarkov.Chain {
	chains := make(map[string]*gomarkov.Chain)
	for _, message := range messages {
		author := message.Author
		if _, ok := chains[author]; !ok {
			chains[author] = gomarkov.NewChain(markovOrder)
		}
		chains[author].Add(strings.Fields(message.Message))
	}
	return chains
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
