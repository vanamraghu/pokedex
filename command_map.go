package main

import (
	"fmt"
	"pokeapis"
)

func displayBackwardLocations(cfg *config) error {
	if cfg.prevUrl == nil {
		return fmt.Errorf("previous url doesn't exist")
	}
	data, err := pokeapis.GetLocationData(cfg.prevUrl)
	if err != nil {
		return err
	}
	cfg.prevUrl = data.PreviousUrl
	cfg.nextUrl = data.NextUrl

	for _, location := range data.Results {
		fmt.Println(location.LocationName)
	}
	return nil
}

func displayLocations(cfg *config) error {
	// if cache has already data, check for the url and display from cache
	data, err := pokeapis.GetLocationData(cfg.nextUrl)
	if err != nil {
		return err
	}
	cfg.nextUrl = data.NextUrl
	cfg.prevUrl = data.PreviousUrl

	for _, location := range data.Results {
		fmt.Println(location.LocationName)
	}
	return nil

}
