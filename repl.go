package main

import(
	"strings"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	fieldsText := strings.Fields(lowerText)
	return fieldsText
}
