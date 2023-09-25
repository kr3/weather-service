package config

type Config struct {
	OpenWeatherAPIURL   string
	OpenWeatherAPIUnits string
	OpenWeatherAPIKey   string
}

// DefaultConfig returns a configuration with the default values.
func DefaultConfig() Config {
	return Config{
		OpenWeatherAPIURL:   "https://api.openweathermap.org/data/2.5/weather",
		OpenWeatherAPIUnits: "imperial",
		OpenWeatherAPIKey:   "XXXXXXXXXXXXXXXXXX", //Replace with real API token
	}
}
