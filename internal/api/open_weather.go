package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"weather-service/internal/config"
)

type OpenWeatherResponse struct {
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
}

func FetchWeatherData(lat string, lon string, config config.Config) (OpenWeatherResponse, error) {
	// Make a request to OpenWeather API
	url := fmt.Sprintf(config.OpenWeatherAPIURL+"?lat=%s&lon=%s&appid=%s&units=%s", lat, lon, config.OpenWeatherAPIKey, config.OpenWeatherAPIUnits)
	log.Printf("Making request: %s", url)

	weatherData := OpenWeatherResponse{}

	resp, err := http.Get(url)
	if err != nil {
		return weatherData, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return weatherData, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	// Parse the API response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return weatherData, err
	}

	if err := json.Unmarshal(body, &weatherData); err != nil {
		return weatherData, err
	}

	return weatherData, nil
}
