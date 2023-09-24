package handlers

import (
	"fmt"
	"net/http"
	"weather-service/internal/api"
	"weather-service/internal/config"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request, config config.Config) {

	// Get latitude and longitude from the request
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	weatherData := api.OpenWeatherResponse{}
	err := error(nil)

	weatherData, err = api.FetchWeatherData(lat, lon, config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Request failed %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// label temp based on ranges
	temperature := ""
	if weatherData.Main.Temperature <= 55 {
		temperature = "Cold"
	} else if weatherData.Main.Temperature <= 80 {
		temperature = "Moderate"
	} else {
		temperature = "Hot"
	}

	alerts := ""
	for _, weather := range weatherData.Weather {
		alerts += weather.Main + " "
	}

	// Prep response
	response := fmt.Sprintf("Temperature: %s\nAlerts: %s", temperature, alerts)

	// Return response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
