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
	filepath := flag.String("file", "", "path to the WhatsApp chat export file")
	flag.Parse()

	if *filepath == "" {
		log.Fatal("Please provide the path to the WhatsApp chat export file using the -file flag")
	}

	messages, err := whatsapp.ParseChat(*filepath)
	if err != nil {
		log.Fatalf("Error parsing chat file: %v", err)
	}

	chain := gomarkov.NewChain(markovOrder)

	for _, message := range messages {
		chain.Add(strings.Fields(message.Message))
	}

	tokens := markov.CreateTokenSlice(chain)

	fmt.Print("Press enter to generate a new message...")
	var generatedMessage string
	for {
		fmt.Scanln()
		attempts := 0
		for {
			generatedMessage = markov.Generate(chain, tokens)
			if !isMessageInList(generatedMessage, messages) {
				break
			}
			attempts++
			if attempts >= retryLimit {
				break
			}
		}
		fmt.Println(generatedMessage)
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
