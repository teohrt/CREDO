package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type WeatherResult struct {
	ZipCode   string `json:"zipcode"`
	Tempature int    `json:"tempature"`
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	arr, err := getFavorites()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(arr)
}

func getFavorites() ([]WeatherResult, error) {
	favorites := []int{50158, 90514, 50011, 904101, 94118}

	result := []WeatherResult{}

	weatherChan := make(chan WeatherResult)
	errChan := make(chan error)

	for _, fav := range favorites {
		go func() {
			weather, err := getWeather(fav)
			if err != nil {
				log.Printf("Error in api call.")

			}

			weatherChan <- weather
		}()
	}

	for i := 0; i < len(favorites); i++ {

		select {
		case resp := <-weatherChan:
			result = append(result, resp)
		case err := <-errChan:
			return result, err
		}
	}

	return result, nil
}

func getWeather(zipCode int) (WeatherResult, error) {
	return WeatherResult{}, nil
}
