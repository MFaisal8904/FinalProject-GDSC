package services

import (
	"encoding/json"
	"fmt"
	"goweather/config"
	"net/http"
)

const (
	baseURL     = "http://api.openweathermap.org/data/2.5/weather"
	units       = "metric"
	contentType = "application/json"
)

// WeatherService provides weather-related services
type WeatherService struct{}

// WeatherResponse represents the weather data returned from OpenWeatherMap API
type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

// GetWeather fetches the weather information for a given city
func (s *WeatherService) GetWeather(city string) (*WeatherResponse, error) {
	apiKey := config.OpenWeatherMapAPIKey
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=%s", baseURL, city, apiKey, units)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching weather data: received status code %d", resp.StatusCode)
	}

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, fmt.Errorf("error decoding weather data: %v", err)
	}

	return &weatherResp, nil
}
