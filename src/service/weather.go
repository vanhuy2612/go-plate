package service

import (
	"encoding/json"
	"io"
	"net/http"
	"root/src/models"
	"time"
)

type IWeatherService interface {
	Fetch()
}

type WeatherService struct {

}

var wClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:       100,
		MaxConnsPerHost:    100,
		IdleConnTimeout:    90 * time.Second,
		DisableCompression: false,
	},
	Timeout: 10 * time.Second,
}

func (w *WeatherService) Fetch() models.Weather {
	resp, err := wClient.Get("https://api.openweathermap.org/data/2.5/onecall?lat=21.1167&lon=105.8833&units=metric&appid=5796abbde9106b7da4febfae8c44c232")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var weather models.Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		panic(err)
	}
	return weather
}


