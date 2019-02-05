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

// APIResult is the struct containing the results from a completed OWM API call.
type APIResult struct {
	ZipCode   string  `json:"zipcode"`
	Tempature float64 `json:"tempature"`
}

// WeatherHandler returns a slice of type WeatherResult containg the zipcodes and current tempatures of our favorite places
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

func getFavorites() ([]APIResult, error) {

	favorites := []int{50158, 94123, 50011, 94108, 94118}

	result := []APIResult{}

	weatherChan := make(chan APIResult)
	errChan := make(chan error)

	// Concurrently send OpenWeatherMap API results over channels
	for _, fav := range favorites {
		go func(f int) {
			weather, err := getWeather(f)
			if err != nil {
				log.Printf("Error in api call.")
				errChan <- err
				return
			}

			weatherChan <- weather
		}(fav)
	}

	// Aggregate weather data from channels
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

func getWeather(zipCode int) (APIResult, error) {
	APIKey := os.Getenv("OWM_DEVELOPER_KEY")
	url := "https://api.openweathermap.org/data/2.5/weather?zip=" + strconv.Itoa(zipCode) + ",us&APPID=" + APIKey

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return APIResult{strconv.Itoa(zipCode), -1}, err
	}

	// []byte <- response.body
	data, _ := ioutil.ReadAll(response.Body)

	// Container for json response
	var jsonInterface interface{}

	if err := json.Unmarshal(data, &jsonInterface); err != nil {
		log.Fatal(err)
	}

	// Type cast everything. Avoids creating a large response object.
	tempature := jsonInterface.(map[string]interface{})["main"].(map[string]interface{})["temp"].(float64)

	return APIResult{strconv.Itoa(zipCode), tempature}, nil
}
