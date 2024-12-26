// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// 	"strings"
// )

// // WeatherResponse represents the initial API response with cities and condition
// type WeatherResponse struct {
// 	City1     string `json:"city1"`
// 	City2     string `json:"city2"`
// 	Condition string `json:"condition"`
// }

// // CityWeather represents the weather details for a city
// type CityWeather struct {
// 	Temperature float64 `json:"temperature"`
// 	Wind       float64 `json:"wind"`
// 	Rain       float64 `json:"rain"`
// 	Cloud      float64 `json:"cloud"`
// }
// func main() {
// 	// Get cities and condition
// 	baseURL := "https://quest.squadcast.tech/api/2021UCH1682/weather"
// 	initialResponse := getInitialData(baseURL)

// 	// Get weather for both cities
// 	city1Weather := getCityWeather(baseURL, initialResponse.City1)
// 	city2Weather := getCityWeather(baseURL, initialResponse.City2)
// 	// Determine better location based on condition
// 	betterLocation := findBetterLocation(
// 		initialResponse.City1,
// 		initialResponse.City2,
// 		city1Weather,
// 		city2Weather,
// 		initialResponse.Condition,
// 	)
// 	// Submit the answer
// 	submitAnswer(baseURL, betterLocation)
// }

// func getInitialData(baseURL string) WeatherResponse {
// 	response, err := http.Get(baseURL)
// 	checkNilErr(err)
// 	defer response.Body.Close()

// 	var weatherData WeatherResponse
// 	body, err := ioutil.ReadAll(response.Body)
// 	checkNilErr(err)
	
// 	err = json.Unmarshal(body, &weatherData)
// 	checkNilErr(err)
	
// 	return weatherData
// }

// func getCityWeather(baseURL string, city string) CityWeather {
// 	cityURL := fmt.Sprintf("%s/get?q=%s", baseURL, url.QueryEscape(city))
// 	response, err := http.Get(cityURL)
// 	checkNilErr(err)
// 	defer response.Body.Close()
// 	var cityWeather CityWeather
// 	body, err := ioutil.ReadAll(response.Body)
// 	checkNilErr(err)
// 	err = json.Unmarshal(body, &cityWeather)
// 	checkNilErr(err)
// 	return cityWeather
// }
// func findBetterLocation(city1, city2 string, weather1, 
// 	weather2 CityWeather, condition string) string {
// 	switch condition {
// 	case "hot":
// 		if weather1.Temperature > weather2.Temperature {
// 			return city1
// 		}
// 		return city2
// 	case "cold":
// 		if weather1.Temperature < weather2.Temperature {
// 			return city1
// 		}
// 		return city2
// 	case "windy":
// 		if weather1.Wind > weather2.Wind {
// 			return city1
// 		}
// 		return city2
// 	case "rainy":
// 		if weather1.Rain > weather2.Rain {
// 			return city1
// 		}
// 		return city2
// 	case "sunny":
// 		if weather1.Cloud < weather2.Cloud {
// 			return city1
// 		}
// 		return city2
// 	case "cloudy":
// 		if weather1.Cloud > weather2.Cloud {
// 			return city1
// 		}
// 		return city2
// 	default:
// 		return city1 // Default case, should not happen given the problem constraints
// 	}
// }

// func submitAnswer(baseURL string, answer string) {
// 	submitURL := fmt.Sprintf("%s/submit/weather?answer=%s&extension=go", 
// 	baseURL, url.QueryEscape(answer))
	
// 	// Get the source code
// 	sourceCode, err := ioutil.ReadFile("main.go")
// 	checkNilErr(err)
	
// 	// Create POST request with source code in body
// 	request, err := http.NewRequest("POST", submitURL, 
// 	strings.NewReader(string(sourceCode)))
// 	checkNilErr(err)
	
// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	checkNilErr(err)
// 	defer response.Body.Close()
	
// 	// Print submission response
// 	content, err := ioutil.ReadAll(response.Body)
// 	checkNilErr(err)
// 	fmt.Println("Submission response:", string(content))
// }

// func checkNilErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// package main

// import (
//     "encoding/json"
//     "fmt"
//     "io/ioutil"
//     "net/http"
//     "net/url"
//     "strings"
// )

// // Struct for initial weather response
// type WeatherData struct {
//     City1     string `json:"city1"`
//     City2     string `json:"city2"`
//     Condition string `json:"condition"`
// }

// // Struct for city weather details
// type CityWeather struct {
//     Temperature float64 `json:"temperature"`
//     Wind       float64 `json:"wind"`
//     Rain       float64 `json:"rain"`
//     Cloud      float64 `json:"cloud"`
// }

// func main() {
//     fmt.Println("Finding better location based on weather conditions")

//     // Get initial data (cities and condition)
//     baseURL := "https://quest.squadcast.tech/api/211113214"
//     weatherData := getInitialData(baseURL + "/weather")

//     // Get weather details for both cities
//     city1Weather := getCityWeather(baseURL+"/weather/get", weatherData.City1)
//     city2Weather := getCityWeather(baseURL+"/weather/get", weatherData.City2)

//     // Find better location
//     betterLocation := findBetterLocation(weatherData, city1Weather, city2Weather)

//     // Submit answer
//     submitAnswer(baseURL+"/submit/weather", betterLocation)
// }

// func getInitialData(url string) WeatherData {
//     response, err := http.Get(url)
//     checkNilErr(err)
//     defer response.Body.Close()

//     body, err := ioutil.ReadAll(response.Body)
//     checkNilErr(err)

//     var weatherData WeatherData
//     err = json.Unmarshal(body, &weatherData)
//     checkNilErr(err)

//     return weatherData
// }

// func getCityWeather(baseURL string, city string) CityWeather {
//     // Create URL with query parameter
//     cityURL := fmt.Sprintf("%s?q=%s", baseURL, url.QueryEscape(city))
    
//     response, err := http.Get(cityURL)
//     checkNilErr(err)
//     defer response.Body.Close()

//     body, err := ioutil.ReadAll(response.Body)
//     checkNilErr(err)

//     var cityWeather CityWeather
//     err = json.Unmarshal(body, &cityWeather)
//     checkNilErr(err)

//     return cityWeather
// }

// func findBetterLocation(data WeatherData, city1Weather, city2Weather CityWeather) string {
//     switch data.Condition {
//     case "hot":
//         if city1Weather.Temperature > city2Weather.Temperature {
//             return data.City1
//         }
//         return data.City2
//     case "cold":
//         if city1Weather.Temperature < city2Weather.Temperature {
//             return data.City1
//         }
//         return data.City2
//     case "windy":
//         if city1Weather.Wind > city2Weather.Wind {
//             return data.City1
//         }
//         return data.City2
//     case "rainy":
//         if city1Weather.Rain > city2Weather.Rain {
//             return data.City1
//         }
//         return data.City2
//     case "sunny":
//         if city1Weather.Cloud < city2Weather.Cloud {
//             return data.City1
//         }
//         return data.City2
//     case "cloudy":
//         if city1Weather.Cloud > city2Weather.Cloud {
//             return data.City1
//         }
//         return data.City2
//     default:
//         return data.City1
//     }
// }

// func submitAnswer(submitURL string, answer string) {
//     // Read the current file's content
//     sourceCode, err := ioutil.ReadFile("main.go")
//     checkNilErr(err)

//     // Create submission URL with parameters
//     fullURL := fmt.Sprintf("%s?answer=%s&extension=go", submitURL, url.QueryEscape(answer))

//     // Create POST request with source code in body
//     request, err := http.NewRequest("POST", fullURL, strings.NewReader(string(sourceCode)))
//     checkNilErr(err)

//     // Make the request
//     client := &http.Client{}
//     response, err := client.Do(request)
//     checkNilErr(err)
//     defer response.Body.Close()

//     // Print submission response
//     content, err := ioutil.ReadAll(response.Body)
//     checkNilErr(err)
//     fmt.Println("Submission response:", string(content))
// }

// func checkNilErr(err error) {
//     if err != nil {
//         panic(err)
//     }
// }




package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
)

// WeatherData represents the initial API response
type WeatherData struct {
    City1     string `json:"city1"`
    City2     string `json:"city2"`
    Condition string `json:"condition"`
}

// CityWeather represents the weather details
type CityWeather struct {
    Temperature float64 `json:"temperature"`
    Wind       float64 `json:"wind"`
    Rain       float64 `json:"rain"`
    Cloud      float64 `json:"cloud"`
}

func main() {
    // Get initial weather data
    weatherData := getWeatherData()
    
    // Get weather details for both cities
    city1Weather := getCityWeather(weatherData.City1)
    city2Weather := getCityWeather(weatherData.City2)
    
    // Find better location
    betterLocation := getBetterLocation(
        weatherData.City1,
        weatherData.City2,
        weatherData.Condition,
        city1Weather,
        city2Weather,
    )
    
    // Submit answer
    submitAnswer(betterLocation)
}

func getWeatherData() WeatherData {
    resp, err := http.Get("https://quest.squadcast.tech/api/211113214/weather")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    var data WeatherData
    if err := json.Unmarshal(body, &data); err != nil {
        panic(err)
    }
    return data
}

func getCityWeather(city string) CityWeather {
    url := fmt.Sprintf("https://quest.squadcast.tech/api/211113214/weather/get?q=%s", 
        url.QueryEscape(city))
    
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    var weather CityWeather
    if err := json.Unmarshal(body, &weather); err != nil {
        panic(err)
    }
    return weather
}

func getBetterLocation(city1, city2, condition string, weather1, weather2 CityWeather) string {
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
        return city1
    }
}

func submitAnswer(answer string) {
    // Read current file content
    sourceCode, err := ioutil.ReadFile("main.go")
    if err != nil {
        panic(err)
    }

    // Create submission URL
    submitURL := fmt.Sprintf("https://quest.squadcast.tech/api/211113214/submit/weather?answer=%s&extension=go",
        url.QueryEscape(answer))

    // Create request with source code in body
    req, err := http.NewRequest("POST", submitURL, strings.NewReader(string(sourceCode)))
    if err != nil {
        panic(err)
    }

    // Make request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    // Print response
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}