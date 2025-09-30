package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type MeteoResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
		WeatherCode int     `json:"weathercode"`
	} `json:"current_weather"`
}

var cityCoords = map[string][2]float64{
	"Moscow":   {55.7558, 37.6173},
	"London":   {51.5074, -0.1278},
	"Paris":    {48.8566, 2.3522},
	"New York": {40.7128, -74.0060},
	"Chicago":  {41.8781, -87.6298},
}

func GetWeather(city string) (string, error) {
	coords, ok := cityCoords[city]
	if !ok {
		return "", fmt.Errorf("unknown city: %s", city)
	}

	lat := fmt.Sprintf("%f", coords[0])
	lon := fmt.Sprintf("%f", coords[1])

	apiURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true", url.QueryEscape(lat), url.QueryEscape(lon))

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var w MeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		return "", err
	}

	return fmt.Sprintf("Weather in %s: %.1f°C, wind %.1f km/h", city, w.CurrentWeather.Temperature, w.CurrentWeather.WindSpeed), nil
}
