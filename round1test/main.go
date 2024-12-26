package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CityWeather struct {
	Temperature float64 `json:"temperature"`
	Wind        struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Rain struct {
		Value *float64 `json:"3h"`
	} `json:"rain"`
	Clouds struct {
		All *float64 `json:"all"`
	} `
	json:"clouds"`
}

const baseURL = "https://quest.squadcast.tech/api/211117074"

func main() {
	city1, city2, condition, err := fetchCityAndCondition()
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	weather1, err := fetchWeather(city1)
	if err != nil {
		log.Fatalf("Error fetching weather for %s: %v", city1, err)
	}

	weather2, err := fetchWeather(city2)
	if err != nil {
		log.Fatalf("Error fetching weather for %s: %v", city2, err)
	}

	result := evaluateCity(city1, city2, weather1, weather2, condition)
	response, err := submitResult(result)
	if err != nil {
		log.Fatalf("Error submitting result: %v", err)
	}

	fmt.Printf("Response: %s\n", response)
}

func fetchCityAndCondition() (string, string, string, error) {
	resp, err := http.Get(baseURL + "/weather")
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", "", "", err
	}

	var city1, city2, condition string
	doc.Find("code").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.Contains(text, "City 1:") {
			city1 = strings.TrimSpace(strings.Split(text, ":")[1])
		} else if strings.Contains(text, "City 2:") {
			city2 = strings.TrimSpace(strings.Split(text, ":")[1])
		} else if strings.Contains(text, "Condition:") {
			condition = strings.TrimSpace(strings.Split(text, ":")[1])
		}
	})

	if city1 == "" || city2 == "" || condition == "" {
		return "", "", "", fmt.Errorf("insufficient data retrieved")
	}

	return city1, city2, condition, nil
}

func fetchWeather(city string) (*CityWeather, error) {
	resp, err := http.Get(baseURL + "/weather/get?q=" + city)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather CityWeather
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

func evaluateCity(city1, city2 string, weather1, weather2 *CityWeather, condition string) string {
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
		if weather1.Wind.Speed > weather2.Wind.Speed {
			return city1
		}
		return city2
	case "rainy":
		if getValue(weather1.Rain.Value) > getValue(weather2.Rain.Value) {
			return city1
		}
		return city2
	case "sunny":
		if getValue(weather1.Clouds.All) < getValue(weather2.Clouds.All) {
			return city1
		}
		return city2
	case "cloudy":
		if getValue(weather1.Clouds.All) > getValue(weather2.Clouds.All) {
			return city1
		}
		return city2
	default:
		return city2
	}
}

func getValue(value *float64) float64 {
	if value == nil {
		return 0.0
	}
	return *value
}

func submitResult(answer string) (string, error) {
	code, err := ioutil.ReadFile("main.go")
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/submit/weather?answer=%s&extension=go", baseURL, answer)
	req, err := http.NewRequest("POST", url, bytes.NewReader(code))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "text/plain")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil
}