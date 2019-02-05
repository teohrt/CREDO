package main

import (
	"log"
	"net/http"

	c "./countcaps"
	w "./weather"
	"github.com/gorilla/mux"
)

func main() {
	initServer()
}

func initServer() {
	r := mux.NewRouter()

	r.HandleFunc("/api/countcaps/{word}", c.CountCapsHandler).Methods("GET")
	r.HandleFunc("/api/weather/favorites", w.WeatherHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
