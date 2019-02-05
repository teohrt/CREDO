package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type WeatherResult struct {
	ZipCode   string  `json:"zipcode"`
	Tempature float64 `json:"tempature"`
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	arr, err := getFavorites()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(arr)
}

func getFavorites() ([]WeatherResult, error) {

	favorites := []int{50158, 90514, 50011, 94108, 94118}

	result := []WeatherResult{}

	weatherChan := make(chan WeatherResult)
	errChan := make(chan error)

	// Concurrently send OpenWeatherMap API results over channels
	for _, fav := range favorites {
		go func() {
			weather, err := getWeather(fav)
			if err != nil {
				log.Printf("Error in api call.")
				errChan <- err
				return
			}

			weatherChan <- weather
		}()
	}

	// Aggregate weather data from channels
	// TODO: there is an error where the last favorite response is the only struct populating the result slice
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
	APIKey := os.Getenv("OWM_DEVELOPER_KEY")
	url := "https://api.openweathermap.org/data/2.5/weather?zip=" + strconv.Itoa(zipCode) + ",us&APPID=" + APIKey

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return WeatherResult{strconv.Itoa(zipCode), -1}, err
	}

	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))

	var jsonInterface interface{}

	if err := json.Unmarshal(data, &jsonInterface); err != nil {
		log.Fatal(err)
	}

	// Type cast everything. Avoids storing the whole json object.
	tempature := jsonInterface.(map[string]interface{})["main"].(map[string]interface{})["temp"].(float64)

	//fmt.Println("testing: ", tempature)

	return WeatherResult{strconv.Itoa(zipCode), tempature}, nil
}
