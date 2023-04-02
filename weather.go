package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
)

type Weather struct {
	Date        string
	WeatherText string
	Temperature string
}

type WeatherReq struct {
	Location string `json:"location"`
	Key      string `json:"key"`
	Start    int    `json:"start"`
	Days     int    `json:"days"`
	Language string `json:"language"`
	Unit     string `json:"unit"`
}

type WeatherResponse struct {
	Results []struct {
		Location struct {
			ID             string `json:"id"`
			Name           string `json:"name"`
			Country        string `json:"country"`
			Path           string `json:"path"`
			Timezone       string `json:"timezone"`
			TimezoneOffset string `json:"timezone_offset"`
		} `json:"location"`
		Daily      `json:"daily"`
		LastUpdate string `json:"last_update"`
	} `json:"results"`
}

type Daily []Day
type Day struct {
	Date                string `json:"date"`
	TextDay             string `json:"text_day"`
	CodeDay             string `json:"code_day"`
	TextNight           string `json:"text_night"`
	CodeNight           string `json:"code_night"`
	High                string `json:"high"`
	Low                 string `json:"low"`
	Precip              string `json:"precip"`
	WindDirection       string `json:"wind_direction"`
	WindDirectionDegree string `json:"wind_direction_degree"`
	WindSpeed           string `json:"wind_speed"`
	WindScale           string `json:"wind_scale"`
	Rainfall            string `json:"rainfall"`
	Humidity            string `json:"humidity"`
}

func getWeather(start, days int) (Daily, error) {
	url := "https://api.seniverse.com/v3/weather/daily.json"
	req := WeatherReq{
		Location: getConfig().SeniverseLocation,
		Key:      getConfig().SeniverseAPIKey,
		Start:    start,
		Days:     days,
		Language: "zh-Hans",
		Unit:     "c",
	}
	var weatherResponse WeatherResponse
	resp, _, errs := gorequest.New().Get(url).Query(req).EndStruct(&weatherResponse)
	if len(errs) > 0 {
		log.Fatalf("Error making request: %v", errs)
		return nil, errs[0]
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Invalid response status: %d", resp.StatusCode)
		return nil, fmt.Errorf("Invalid response status: %d", resp.StatusCode)
	}

	return weatherResponse.Results[0].Daily, nil
}
