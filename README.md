# Weather Application

This is a Go application that provides weather information based on latitude and longitude coordinates. It uses the OpenWeather API to fetch weather data and returns the temperature and weather alerts for the specified location.

## Prerequisites

Before running this application, you need to obtain an API key from OpenWeather. You can sign up for a free API key on the [OpenWeather website](https://openweathermap.org/api) and replace the `OpenWeatherAPIKey` field in the DefaultConfig with your own API key.

## Usage

To retrieve weather information, make a GET request to the `/weather` endpoint with latitude and longitude query parameters. For example:

```bash
http://localhost:8080/weather?lat=37.7749&lon=-122.4194
```

This will return a response in plain text format with the current temperature and weather alerts for the specified location.

## Response Format

The response will have the following format:

```
Temperature: <Temperature>
Alerts: <Weather Alerts>
```

- `<Temperature>` will be one of the following: "Cold" "Moderate" or "Hot" depending on the temperature.
- `<Weather Alerts>` will include a list of weather descriptions.

## Acknowledgments

- This application uses the OpenWeather API to fetch weather data.
- Built with Go (Golang).