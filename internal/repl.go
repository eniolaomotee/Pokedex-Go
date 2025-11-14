package internal

import (
	"strings"
)

func cleanInput(text string) []string{
	cleanText := strings.ToLower(text)
	words := strings.Fields(cleanText)
	return words

}