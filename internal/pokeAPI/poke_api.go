package pokeAPI

import (
	"encoding/json"
	"net/http"
	"fmt"
)

const (
	location_area_endpoint = "https://pokeapi.co/api/v2/location-area/"
)

type LocationAreasAPIResults struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	LocationAreas  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
}

func GetLocationAreas(parameters string) ([]string, error) {
	fullURL := location_area_endpoint
	res, err := http.Get(fullURL)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var decodedResults LocationAreasAPIResults 
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&decodedResults)
	if err != nil {
		return nil, err
	}

	locations := []string{}
	for i := range decodedResults.LocationAreas {
		locations = append(locations, decodedResults.LocationAreas[i].Name)
	}
	fmt.Println("DEBUG:", locations)
	return locations, nil
}