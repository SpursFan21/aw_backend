package main

import (
	"encoding/json"
	"net/http"
)

func getWeather(w http.ResponseWriter, r *http.Request) {
	apiKey := "AX0VBcCNEgbPL21wTdOonbtcfKynYTH4"
	url := "https://api.tomorrow.io/v4/timelines?location=-36.8485,174.7633&fields=temperature,precipitationIntensity&timesteps=1h&units=metric&apikey=" + apiKey

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var weatherData map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(weatherData)
}

func main() {
	http.HandleFunc("/api/weather", getWeather)
	http.ListenAndServe(":8080", nil)
}
