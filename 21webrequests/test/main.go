package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
// [Previous struct definitions remain the same]
type WeatherResponse struct {
	City1     string `json:"city1"`
	City2     string `json:"city2"`
	Condition string `json:"condition"`
}
type CityWeather struct {
	Temperature float64 `json:"temperature"`
	Wind       float64 `json:"wind"`
	Rain       float64 `json:"rain"`
	Cloud      float64 `json:"cloud"`
}
func main() {
	baseURL := "https://quest.squadcast.tech/api/2021UCH1682/weather"
	// Add debug output
	fmt.Printf("Making request to: %s\n", baseURL)
	response, err := http.Get(baseURL)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer response.Body.Close()
	// Print response status
	fmt.Printf("Response status: %s\n", response.Status)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read body: %v", err)
	}
	// Print raw response
	fmt.Printf("Raw response: %s\n", string(body))
	var weatherData WeatherResponse
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
	if weatherData.City1 == "" || weatherData.City2 == "" || weatherData.Condition == "" {
		log.Fatal("Invalid response: missing required fields")
	}
	// Continue with rest of program...
	// [Rest of the functions remain the same]
}
func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err) // Changed from panic to log.Fatal for better error messages
	}
}