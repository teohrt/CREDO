package weather

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

// Obviously incomplete testing. Testing handlers is quite comprehensive
func TestGetWeather(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	zipCode := 50158
	correct := WeatherResult{"50158", 260.13}

	result, err := getWeather(zipCode)
	if err != nil {
		t.Errorf("API error")
		return
	}

	if result != correct {
		t.Error("Incorrect result. Expected 50158, -1. Got: ", result)
	}
}
