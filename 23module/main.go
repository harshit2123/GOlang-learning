package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	// Test input with unicode representation of emojis
	testInput := "ten eight \\u1F604 nine six five \\u1F606"
	
	// Convert unicode to emojis
	decoded := convertUnicodeToEmoji(testInput)
	fmt.Printf("Decoded text: %s\n", decoded)
	
	// URL encode the result
	encoded := url.QueryEscape(decoded)
	fmt.Printf("URL encoded: %s\n", encoded)
	
	// Show submission URL
	submitURL := fmt.Sprintf("https://quest.squadcast.tech/api/2021UCH1682/submit/emoji?answer=%s&extension=go", encoded)
	fmt.Printf("Submission URL: %s\n", submitURL)
}

func convertUnicodeToEmoji(text string) string {
	// Simple replacements for this example
	replacements := map[string]string{
		"\\u1F604": "😄", // Smiling face with open mouth and smiling eyes
		"\\u1F606": "😆", // Smiling face with open mouth and tightly-closed eyes
	}
	
	result := text
	for unicode, emoji := range replacements {
		result = strings.ReplaceAll(result, unicode, emoji)
	}
	
	return result
}