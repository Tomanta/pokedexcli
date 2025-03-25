package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args []string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.previousLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.LocationAreas {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you are on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.previousLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.LocationAreas {
		fmt.Println(location.Name)
	}
	return nil
}
