package markov

import (
	"strings"

	"github.com/mb-14/gomarkov"
)

// Generate generates a new message based on the Markov chain
func Generate(chain *gomarkov.Chain, tokens []string) string {
	order := chain.Order

	// Generate a new message until the end token is generated
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		// Get the last n tokens where n is the order of the chain
		recentTokens := tokens[(len(tokens) - order):]
		next, _ := chain.Generate(recentTokens)
		tokens = append(tokens, next)
	}
	generatedMessage := strings.Join(tokens[order:len(tokens)-1], " ")

	return generatedMessage
}
