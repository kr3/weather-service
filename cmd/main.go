package main

import (
	"net/http"
	"weather-service/internal/config"
	"weather-service/internal/handlers"
)

func main() {
	config := config.DefaultConfig()
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		handlers.WeatherHandler(w, r, config)
	})
	http.ListenAndServe(":8080", nil)
}
