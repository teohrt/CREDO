package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initServer()
}

func initServer() {
	r := mux.NewRouter()

	r.HandleFunc("/api/countcaps/{word}", CountCapsHandler).Methods("GET")
	r.HandleFunc("/api/weather/favorites", WeatherHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
