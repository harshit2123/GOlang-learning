// package main
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// 	"strings"
// )
// var numberWords = map[string]string{
// 	"zero":  "0",
// 	"one":   "1",
// 	"two":   "2",
// 	"three": "3",
// 	"four":  "4",
// 	"five":  "5",
// 	"six":   "6",
// 	"seven": "7",
// 	"eight": "8",
// 	"nine":  "9",
// 	"point": ".",
// }
// func main() {
// 	// Get the passage
// 	baseURL := "https://quest.squadcast.tech/api/2021UCH1682/worded_ip"
// 	passage := getPassage(baseURL)
// 	// Find and convert IP
// 	ipAddress := findIP(passage)
// 	// Submit answer
// 	submitAnswer(baseURL, ipAddress)
// }
// func getPassage(baseURL string) string {
// 	response, err := http.Get(baseURL)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()

// 	content, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(content)
// }
// func findIP(passage string) string {
// 	words := strings.Fields(passage)
// 	var ipParts []string
// 	var result strings.Builder
// 	for _, word := range words {
// 		word = strings.ToLower(word)
// 		if num, exists := numberWords[word]; exists {
// 			if num == "." {
// 				if result.Len() > 0 {
// 					ipParts = append(ipParts, result.String())
// 					result.Reset()
// 				}
// 			} else {
// 				result.WriteString(num)
// 			}}}
// 	// Add the last part if exists
// 	if result.Len() > 0 {
// 		ipParts = append(ipParts, result.String())
// 	}
// 	// Validate IP parts
// 	if len(ipParts) != 4 {
// 		panic("Invalid IP format found")
// 	}
// 	return strings.Join(ipParts, ".")
// }
// func submitAnswer(baseURL string, answer string) {
// 	submitURL := fmt.Sprintf("%s/submit/worded_ip?answer=%s&extension=go",
// 		strings.TrimSuffix(baseURL, "/worded_ip"),
// 		url.QueryEscape(answer))
// 	sourceCode, err := ioutil.ReadFile("main.go")
// 	if err != nil {
// 		panic(err)
// 	}
// 	request, err := http.NewRequest("POST", submitURL, strings.NewReader(string(sourceCode)))
// 	if err != nil {
// 		panic(err)
// 	}
// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()
// 	content, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Submission response:", string(content))
// }

// package main
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// 	"strconv"
// 	"strings"
// 	"github.com/fernet/fernet-go"
// )
// type Response struct {
// 	EncryptedText string `json:"encrypted_text"`
// 	Key           string `json:"key"`
// }
// func main() {
// 	baseURL := "https://quest.squadcast.tech/api/2021UCH1682/emoji"
// 	// Get encrypted data
// 	response := getEncryptedData(baseURL)
// 	// Decrypt text
// 	decryptedText := decryptText(response.EncryptedText, response.Key)
// 	// Convert unicode to emojis
// 	finalText := convertUnicodeToEmoji(decryptedText)
// 	// Submit answer
// 	submitAnswer(baseURL, finalText)
// }
// func getEncryptedData(baseURL string) Response {
// 	resp, err := http.Get(baseURL)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var response Response
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return response
// }
// func decryptText(encrypted, key string) string {
// 	k, err := fernet.DecodeKey(key)
// 	if err != nil {
// 		panic(err)
// 	}
// 	tok := fernet.VerifyAndDecrypt([]byte(encrypted), 0, []*fernet.Key{k})
// 	if tok == nil {
// 		panic("Failed to decrypt")
// 	}
// 	return string(tok)
// }
// func convertUnicodeToEmoji(text string) string {
// 	// Convert unicode escape sequences to actual emojis
// 	result := text
// 	// Find sequences like \uXXXX or U+XXXX and convert them
// 	// This is a simplified example - you might need to handle more formats
// 	unicodeReplacer := strings.NewReplacer(
// 		"\\u", "\\u",
// 		"U+", "\\u",
// 	)
// 	result = unicodeReplacer.Replace(result)
// 	// Unescape the unicode
// 	result = unquoteString(result)
// 	return result
// }
// func unquoteString(s string) string {
// 	// Convert escaped unicode to actual characters
// 	unquoted, err :=
// 	strconv.Unquote(`"` + strings.Replace(s, `"`, `\"`, -1) + `"`)
// 	if err != nil {
// 		return s
// 	}
// 	return unquoted
// }
// func submitAnswer(baseURL string, answer string) {
// 	submitURL :=
// 	fmt.Sprintf("%s/submit/emoji?answer=%s&extension=go",
// 		strings.TrimSuffix(baseURL, "/emoji"),
// 		url.QueryEscape(answer))

// 	sourceCode, err := ioutil.ReadFile("main.go")
// 	if err != nil {
// 		panic(err)
// 	}

// 	request, err :=
// 	http.NewRequest("POST", submitURL,
// 	strings.NewReader(string(sourceCode)))
// 	if err != nil {
// 		panic(err)
// 	}

// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()

// 	content, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Submission response:", string(content))
// }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/fernet/fernet-go"
)

type Response struct {
	EncryptedText string `json:"encrypted_text"`
	Key           string `json:"key"`
}

func getEncryptedData(baseURL string) (Response, error) {
	resp, err := http.Get(baseURL)
	if err != nil {
		return Response{}, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK (200)
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body) // Read the body for debugging
		return Response{}, fmt.Errorf("received status code %d. Response body: %s", resp.StatusCode, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, fmt.Errorf("failed to read response body: %w", err)
	}

	// Print the response body for debugging purposes
	fmt.Println("Response Body:", string(body))

	// Now try unmarshaling the response
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}

func decryptText(encrypted, key string) (string, error) {
	k, err := fernet.DecodeKey(key)
	if err != nil {
		return "", fmt.Errorf("failed to decode key: %w", err)
	}

	tok := fernet.VerifyAndDecrypt([]byte(encrypted), 0, []*fernet.Key{k})
	if tok == nil {
		return "", fmt.Errorf("failed to decrypt text")
	}

	return string(tok), nil
}

func convertUnicodeToEmoji(text string) string {
	re := regexp.MustCompile(`(\\u[0-9a-fA-F]{4}|U\+[0-9a-fA-F]{4})`)
	
	return re.ReplaceAllStringFunc(text, func(match string) string {
		if strings.HasPrefix(match, "U+") {
			match = "\\u" + match[2:]
		}
		
		s := fmt.Sprintf("\"%s\"", match)
		unquoted, err := strconv.Unquote(s)
		if err != nil {
			return match
		}
		return unquoted
	})
}

func submitAnswer(baseURL, answer string) error {
	submitURL := fmt.Sprintf("%s/submit/emoji?answer=%s&extension=go", baseURL, url.QueryEscape(answer))
	
	sourceCode, err := ioutil.ReadFile("main.go")
	if err != nil {
		return fmt.Errorf("failed to read source code: %w", err)
	}

	request, err := http.NewRequest("POST", submitURL, strings.NewReader(string(sourceCode)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	fmt.Println("Submission response:", string(content))
	return nil
}

func main() {
	baseURL := "https://quest.squadcast.tech/api/211113214/emoji"
	
	response, err := getEncryptedData(baseURL)
	if err != nil {
		fmt.Printf("Error getting encrypted data: %v\n", err)
		return
	}

	decryptedText, err := decryptText(response.EncryptedText, response.Key)
	if err != nil {
		fmt.Printf("Error decrypting text: %v\n", err)
		return
	}

	finalText := convertUnicodeToEmoji(decryptedText)

	err = submitAnswer(baseURL, finalText)
	if err != nil {
		fmt.Printf("Error submitting answer: %v\n", err)
		return
	}
}