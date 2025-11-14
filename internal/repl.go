package internal

import (
	"strings"
)

func CleanInput(text string) []string{
	cleanText := strings.ToLower(text)
	words := strings.Fields(cleanText)
	return words
}

