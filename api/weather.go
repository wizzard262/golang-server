package handler

import (
	//
	"encoding/json" // encoding and decoding JSON data
	"net/http"      // HTTP server + client (handlers, ResponseWriter, Request)
	"net/url"       // URL parsing and encoding
	"strconv"       // String conversion
)

// OpenMeteoResponse represents the structure of the response from the Open-Meteo API.
// DTO: This struct is used to unmarshal the JSON response from the Open-Meteo API into a Go data structure.
type OpenMeteoResponse struct {
	Current struct {
		Temperature2m float64 `json:"temperature_2m"`
	} `json:"current"`
}

// DTO: Output to be passed out of the function
type OutputResponse struct {
	Lat           float64 `json:"latitude"`
	Lng           float64 `json:"longitude"`
	Temperature2m float64 `json:"temperature_2m"`
}

func Weather(responseWriter http.ResponseWriter, request *http.Request) {
	// Read lat/lon from query params
	lat := request.URL.Query().Get("lat")
	lng := request.URL.Query().Get("lon")

	endpointUrl := "https://api.open-meteo.com/v1/forecast"

	querystringValues := url.Values{}
	querystringValues.Set("latitude", lat)
	querystringValues.Set("longitude", lng)
	querystringValues.Set("current", "temperature_2m")

	resp, err := http.Get(endpointUrl + "?" + querystringValues.Encode())

	// Check for errors in the HTTP request to the Open-Meteo API
	if err != nil {
		http.Error(responseWriter, "failed to call Open-Meteo", http.StatusBadGateway)
		return
	}

	// ensure the HTTP response body is closed after we're done reading it
	defer resp.Body.Close()

	// Populate the OpenMeteoResponse struct with the openMeteoResponse from the Open-Meteo API response
	var openMeteoResponse OpenMeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&openMeteoResponse); err != nil {
		http.Error(responseWriter, "failed to decode Open-Meteo response", http.StatusInternalServerError)
		return
	}

	//
	latFloat, _ := strconv.ParseFloat(lat, 64)
	lngFloat, _ := strconv.ParseFloat(lng, 64)

	// Map to
	var outputResponse OutputResponse
	outputResponse.Lat = latFloat
	outputResponse.Lng = lngFloat
	outputResponse.Temperature2m = openMeteoResponse.Current.Temperature2m

	// Set the response header to indicate that we're returning JSON data
	responseWriter.Header().Set("Content-Type", "application/json")

	// Encode the OutputResponse struct as JSON and write it to the response
	json.NewEncoder(responseWriter).Encode(outputResponse)
}

// Note: as we are hosting on Vercel and it will handle the HTTP server we dont
// include things like: http.Handle() and http.ListenAndServe()
// Vercel only runs serverless functions, not a Go web server.
