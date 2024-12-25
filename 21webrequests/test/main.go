package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// WeatherResponse represents the initial API response with cities and condition
type WeatherResponse struct {
	City1     string `json:"city1"`
	City2     string `json:"city2"`
	Condition string `json:"condition"`
}

// CityWeather represents the weather details for a city
type CityWeather struct {
	Temperature float64 `json:"temperature"`
	Wind       float64 `json:"wind"`
	Rain       float64 `json:"rain"`
	Cloud      float64 `json:"cloud"`
}

func main() {
	// Get cities and condition
	baseURL := "https://quest.squadcast.tech/api/2021UCH1682/weather"
	initialResponse := getInitialData(baseURL)

	// Get weather for both cities
	city1Weather := getCityWeather(baseURL, initialResponse.City1)
	city2Weather := getCityWeather(baseURL, initialResponse.City2)

	// Determine better location based on condition
	betterLocation := findBetterLocation(
		initialResponse.City1,
		initialResponse.City2,
		city1Weather,
		city2Weather,
		initialResponse.Condition,
	)

	// Submit the answer
	submitAnswer(baseURL, betterLocation)
}

func getInitialData(baseURL string) WeatherResponse {
	response, err := http.Get(baseURL)
	checkNilErr(err)
	defer response.Body.Close()

	var weatherData WeatherResponse
	body, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	
	err = json.Unmarshal(body, &weatherData)
	checkNilErr(err)
	
	return weatherData
}

func getCityWeather(baseURL string, city string) CityWeather {
	cityURL := fmt.Sprintf("%s/get?q=%s", baseURL, url.QueryEscape(city))
	response, err := http.Get(cityURL)
	checkNilErr(err)
	defer response.Body.Close()

	var cityWeather CityWeather
	body, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	
	err = json.Unmarshal(body, &cityWeather)
	checkNilErr(err)
	
	return cityWeather
}

func findBetterLocation(city1, city2 string, weather1, weather2 CityWeather, condition string) string {
	switch condition {
	case "hot":
		if weather1.Temperature > weather2.Temperature {
			return city1
		}
		return city2
	case "cold":
		if weather1.Temperature < weather2.Temperature {
			return city1
		}
		return city2
	case "windy":
		if weather1.Wind > weather2.Wind {
			return city1
		}
		return city2
	case "rainy":
		if weather1.Rain > weather2.Rain {
			return city1
		}
		return city2
	case "sunny":
		if weather1.Cloud < weather2.Cloud {
			return city1
		}
		return city2
	case "cloudy":
		if weather1.Cloud > weather2.Cloud {
			return city1
		}
		return city2
	default:
		return city1 // Default case, should not happen given the problem constraints
	}
}

func submitAnswer(baseURL string, answer string) {
	submitURL := fmt.Sprintf("%s/submit/weather?answer=%s&extension=go", baseURL, url.QueryEscape(answer))
	
	// Get the source code
	sourceCode, err := ioutil.ReadFile("main.go")
	checkNilErr(err)
	
	// Create POST request with source code in body
	request, err := http.NewRequest("POST", submitURL, strings.NewReader(string(sourceCode)))
	checkNilErr(err)
	
	client := &http.Client{}
	response, err := client.Do(request)
	checkNilErr(err)
	defer response.Body.Close()
	
	// Print submission response
	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println("Submission response:", string(content))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}